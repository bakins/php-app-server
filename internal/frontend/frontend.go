package frontend

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/gorilla/mux"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"

	"github.com/bakins/php-app-server/internal/fastcgi"
	httpserver "github.com/bakins/php-app-server/internal/http"
)

type Config struct {
	Target   *url.URL `kong:"required"`
	Address  string   `kong:"default=:8080"`
	Filename string   `kong:"required"`
}

type Server struct {
	*httpserver.Server
}

func FromConfig(cfg Config) (*Server, error) {
	server, err := httpserver.NewServer(
		httpserver.WithServerAddress("tcp", cfg.Address),
	)
	if err != nil {
		return nil, err
	}

	f, err := fastcgi.New(cfg.Target.Scheme, cfg.Target.Host, cfg.Filename)
	if err != nil {
		return nil, err
	}

	// keep any garbage requests from forwarding to php
	matcher := func(r *http.Request, _ *mux.RouteMatch) bool {
		ct := r.Header.Get("Content-Type")

		return strings.HasPrefix(ct, "application/json") ||
			strings.HasPrefix(ct, "application/protobuf")
	}

	// twirp is POST only
	server.Handle("/twirp/", f).Methods("POST").MatcherFunc(matcher)

	server.AddMiddleware(func(next http.Handler) http.Handler {
		return otelhttp.NewHandler(next, "frontend")
	})

	s := Server{
		Server: server,
	}

	return &s, nil
}
