package fastcgi

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"path/filepath"

	"github.com/yookoala/gofast"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/propagation"
	"go.uber.org/zap"
)

// Handler is a http.Handler that proxies using fastcgi to a single endpoint.
type Handler struct {
	handler http.Handler
	logger  *zap.Logger
}

// New creates a handler.
//
// filename does not have to exist - it could be in another
// container in same pod. This filename is as the fastcgi process
// sees the filename, not this process.
func New(network string, address string, filename string, options ...Option) (*Handler, error) {
	opts := config{
		logger: zap.NewNop(),
	}

	for _, o := range options {
		o.apply(&opts)
	}

	abs, err := filepath.Abs(filename)
	if err != nil {
		return nil, err
	}

	mw := []gofast.Middleware{
		gofast.BasicParamsMap,
		gofast.MapHeader,
		gofast.MapEndpoint(abs),
	}

	mw = append(mw, opts.middleware...)

	chain := gofast.Chain(mw...)

	connFactory := gofast.SimpleConnFactory(network, address)

	f := fastcgiTransport{
		newClient:      gofast.SimpleClientFactory(connFactory),
		sessionHandler: chain(gofast.BasicSession),
	}

	propogator := propagation.TraceContext{}

	proxy := httputil.ReverseProxy{
		Transport: otelhttp.NewTransport(
			&f,
			otelhttp.WithPropagators(propogator),
			otelhttp.WithSpanNameFormatter(func(operation string, r *http.Request) string {
				return "fastcgi"
			}),
		),
		Director: func(r *http.Request) {},
	}

	h := Handler{
		handler: &proxy,
		logger:  opts.logger,
	}

	return &h, nil
}

/*
func middleware(next http.Handler) http.Handler {
	propogator := propagation.TraceContext{}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tr := otel.GetTracerProvider().Tracer("php-app-server")

		ctx, span := tr.Start(r.Context(), "fastcgi", trace.WithSpanKind(trace.SpanKindClient))
		defer span.End()

		propogator.Inject(ctx, propagation.HeaderCarrier(r.Header))

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}*/

type fastcgiTransport struct {
	sessionHandler gofast.SessionHandler
	newClient      gofast.ClientFactory
}

func (f *fastcgiTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	c, err := f.newClient()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to FastCGI application %w", err)
	}

	defer c.Close()

	pipe, err := f.sessionHandler(c, gofast.NewRequest(r))
	if err != nil {
		return nil, fmt.Errorf("failed to process FastCGI request %w", err)
	}

	defer pipe.Close()

	var errBuffer bytes.Buffer

	w := &fastcgiResponse{}

	if err = pipe.WriteTo(w, &errBuffer); err != nil {
		return nil, fmt.Errorf("failed writing error buffer to response %w", err)
	}

	// log this?
	//if errBuffer.Len() > 0 {
	//	log.Printf("gofast: error stream from application process %s",
	//		errBuffer.String())
	//}

	resp := http.Response{
		StatusCode: w.statusCode,
		Header:     w.headers,
		Body:       ioutil.NopCloser(&w.Buffer),
	}

	if w.statusCode == 0 {
		w.statusCode = http.StatusOK
	}

	return &resp, nil
}

type fastcgiResponse struct {
	headers http.Header
	bytes.Buffer
	statusCode int
}

func (f *fastcgiResponse) Header() http.Header {
	if f.headers == nil {
		f.headers = make(http.Header, 3)
	}

	return f.headers
}

func (f *fastcgiResponse) WriteHeader(statusCode int) {
	if statusCode == 0 {
		statusCode = http.StatusOK
	}

	f.statusCode = statusCode
}

var _ http.ResponseWriter = &fastcgiResponse{}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.handler.ServeHTTP(w, r)
}

// Option is used when creating Handler
type Option interface {
	apply(*config)
}

type config struct {
	logger     *zap.Logger
	middleware []gofast.Middleware
}

type optionFunc func(*config)

func (f optionFunc) apply(o *config) {
	f(o)
}

// WithMiddleware adds middleware to the Handler. This will reset any middleware
// already set with WithHandlerMiddleware.
// This middleware is added after the base middleware.
func WithMiddleware(middleware ...gofast.Middleware) Option {
	return optionFunc(func(o *config) {
		o.middleware = middleware
	})
}

func WithLogger(logger *zap.Logger) Option {
	return optionFunc(func(o *config) {
		o.logger = logger
	})
}
