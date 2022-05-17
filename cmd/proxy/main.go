package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/alecthomas/kong"

	"github.com/bakins/php-app-server/internal/logging"
	"github.com/bakins/php-app-server/internal/proxy"
)

type config struct {
	Address string `kong:"default=':18080'"`
}

func main() {
	var cfg config

	kong.Parse(&cfg)

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGQUIT)
	defer cancel()

	if err := run(ctx, cfg); err != nil {
		logging.Exit(err)
	}
}

func run(ctx context.Context, cfg config) error {
	logger := logging.NewLogger()

	p, err := proxy.New(
		proxy.WithLogger(logger),
		proxy.WithAddress("tcp", cfg.Address),
	)
	if err != nil {
		return err
	}

	if err := p.Run(ctx); err != nil {
		return logging.NewLoggingError(logger, "proxy failed", err)
	}

	return nil
}
