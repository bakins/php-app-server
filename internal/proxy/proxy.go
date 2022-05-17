package proxy

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"sync"
	"time"

	"github.com/justinas/alice"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

type config struct {
	transport http.RoundTripper
	logger    *zap.Logger
	chain     alice.Chain
	network   string
	address   string
}

type Option interface {
	apply(*config) error
}

type optionFunc func(*config) error

func (f optionFunc) apply(c *config) error {
	return f(c)
}

func WithTransport(transport http.RoundTripper) Option {
	return optionFunc(func(c *config) error {
		if transport == nil {
			return errors.New("transport must not be nil")
		}

		c.transport = transport

		return nil
	})
}

func WithLogger(logger *zap.Logger) Option {
	return optionFunc(func(c *config) error {
		if logger == nil {
			return errors.New("logger must not be nil")
		}

		c.logger = logger

		return nil
	})
}

func WithMiddlewareChain(chain alice.Chain) Option {
	return optionFunc(func(c *config) error {
		c.chain = chain
		return nil
	})
}

func WithAddress(network string, address string) Option {
	return optionFunc(func(c *config) error {
		switch network {
		case "udp", "unixgram":
			return fmt.Errorf("unsupported network %q", network)
		case "":
			return errors.New("network must not be empty")
		}

		if address == "" {
			return errors.New("address must not be empty")
		}

		c.network = network
		c.address = address

		return nil
	})
}

type Server struct {
	handler  http.Handler
	listener net.Listener
	network  string
	address  string
}

func New(options ...Option) (*Server, error) {
	cfg := config{
		transport: http.DefaultTransport,
		logger:    zap.NewNop(),
		network:   "tcp",
		address:   "127.0.0.1:0",
	}

	for _, o := range options {
		if err := o.apply(&cfg); err != nil {
			return nil, err
		}
	}

	h, err := newReverseProxy(cfg.transport, cfg.logger)
	if err != nil {
		return nil, err
	}

	s := Server{
		handler: cfg.chain.Then(h),
		network: cfg.network,
		address: cfg.address,
	}

	return &s, nil
}

func (s *Server) Run(ctx context.Context) error {
	listener, err := net.Listen(s.network, s.address)
	if err != nil {
		return fmt.Errorf(
			"failed to listen %q %q %w",
			s.network,
			s.address,
			err,
		)
	}

	// technically a race condition
	s.listener = listener

	svr := &http.Server{
		Handler: http.HandlerFunc(s.serveHTTP),
	}

	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		if err := svr.Serve(listener); err != nil {
			if err != http.ErrServerClosed {
				return err
			}
		}

		return nil
	})

	eg.Go(func() error {
		<-ctx.Done()
		// allow adjusting timeout?
		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), time.Second*10)
		defer shutdownCancel()

		_ = svr.Shutdown(shutdownCtx)

		return nil
	})

	return eg.Wait()
}

func (s *Server) serveHTTP(w http.ResponseWriter, r *http.Request) {
	if s == nil || s.handler == nil {
		http.Error(w, "server not correctly configured", http.StatusInternalServerError)
		return
	}

	if !r.URL.IsAbs() {
		http.Error(w, "not a valid proxy url", http.StatusBadRequest)
		return
	}

	if r.Method == http.MethodConnect {
		http.Error(w, "CONNECT is not allowed", http.StatusMethodNotAllowed)
		return
	}

	s.handler.ServeHTTP(w, r)
}

const SchemeHeader = "X-Proxy-Scheme"

func newReverseProxy(transport http.RoundTripper, logger *zap.Logger) (*httputil.ReverseProxy, error) {
	proxyLog, err := zap.NewStdLogAt(logger, zap.WarnLevel)
	if err != nil {
		return nil, fmt.Errorf("failed to create a standard logger %w", err)
	}

	p := httputil.ReverseProxy{
		Director: func(r *http.Request) {
			// this allows us to handle https requests without using CONNECT.
			if scheme := r.Header.Get(SchemeHeader); scheme != "" {
				r.Header.Del(SchemeHeader)
				r.URL.Scheme = scheme
			}
		},
		BufferPool: newBufferPool(),
		ErrorLog:   proxyLog,
		ErrorHandler: func(w http.ResponseWriter, r *http.Request, err error) {
			logger.Warn("proxy error", zap.Error(err), zap.String("target", r.URL.String()))
			http.Error(w, err.Error(), http.StatusBadGateway)
		},
		Transport: transport,
	}

	return &p, nil
}

type bufferPool struct {
	sync.Pool
}

func newBufferPool() *bufferPool {
	b := bufferPool{
		Pool: sync.Pool{
			New: func() interface{} {
				b := make([]byte, 8*1024)
				return &b
			},
		},
	}

	return &b
}

func (b *bufferPool) Get() []byte {
	buf := b.Pool.Get().(*[]byte)
	return *buf
}

func (b *bufferPool) Put(buf []byte) {
	b.Pool.Put(&buf)
}

// WaitForAddress waits until an address is assigned. Useful when using
// a random port.
func (s *Server) WaitForAddress(ctx context.Context) (net.Addr, error) {
	// so cheesy
	t := time.NewTicker(time.Millisecond * 100)
	defer t.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-t.C:
			if s.listener != nil {
				return s.listener.Addr(), nil
			}
		}
	}
}
