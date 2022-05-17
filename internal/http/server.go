package http

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/NYTimes/gziphandler"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"golang.org/x/sync/errgroup"
)

// Server is an HTTP server
type Server struct {
	chain    alice.Chain
	listener net.Listener
	router   *mux.Router
	network  string
	address  string
}

type serverConfig struct {
	network string
	address string
}

type ServerOption interface {
	apply(*serverConfig) error
}

type serverOptionFunc func(*serverConfig) error

func (f serverOptionFunc) apply(c *serverConfig) error {
	return f(c)
}

func WithServerAddress(network string, address string) ServerOption {
	return serverOptionFunc(func(c *serverConfig) error {
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

// New creates a new HTTP server
func NewServer(options ...ServerOption) (*Server, error) {
	cfg := serverConfig{
		network: "tcp",
		address: "127.0.0.1:0",
	}

	for _, o := range options {
		if err := o.apply(&cfg); err != nil {
			return nil, fmt.Errorf("failed to create HTTP server %w", err)
		}
	}

	s := &Server{
		network: cfg.network,
		address: cfg.address,
		router:  mux.NewRouter(),
	}

	s.AddMiddleware(func(next http.Handler) http.Handler {
		return h2c.NewHandler(next, &http2.Server{})
	})

	s.AddMiddleware(gziphandler.GzipHandler)

	return s, nil
}

// Handle adds a handler for the given path.
func (s *Server) Handle(path string, handler http.Handler) *mux.Route {
	return s.router.PathPrefix(path).Handler(handler)
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
		Handler: s.chain.Then(s.router),
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

// TODO: allow setting a matcher on middleware?
func (s *Server) AddMiddleware(middleware func(http.Handler) http.Handler) {
	s.chain = s.chain.Append(middleware)
}

// WaitForAddress waits until an address is assigned. Useful when generating
// a listening socket.
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
