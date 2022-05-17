package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/alecthomas/kong"
	"go.uber.org/zap"

	"github.com/bakins/php-app-server/internal/frontend"
	"github.com/bakins/php-app-server/internal/logging"
	"github.com/bakins/php-app-server/internal/otel"
)

type config struct {
	frontend.Config
	otel.TraceConfig   `kong:"embed,prefix=trace."`
	otel.MetricsConfig `kong:"embed,prefix=metrics."`
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
	logger := logging.NewLogger().With(zap.String("component", "frontend"))

	cleanup, err := otel.Create(
		ctx,
		cfg.TraceConfig,
		cfg.MetricsConfig,
	)
	if err != nil {
		return err
	}

	defer cleanup()

	f, err := frontend.FromConfig(cfg.Config)
	if err != nil {
		return err
	}

	if err := f.Run(ctx); err != nil {
		return logging.NewLoggingError(logger, "http server failed", err)
	}

	return nil
}
