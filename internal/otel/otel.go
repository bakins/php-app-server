package otel

import (
	"context"
	"fmt"
	"time"

	gcppropagator "github.com/GoogleCloudPlatform/opentelemetry-operations-go/propagator"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/propagation"
	controller "go.opentelemetry.io/otel/sdk/metric/controller/basic"
	processor "go.opentelemetry.io/otel/sdk/metric/processor/basic"
	"go.opentelemetry.io/otel/sdk/metric/selector/simple"
	"go.opentelemetry.io/otel/sdk/trace"
)

type TraceConfig struct {
	Endpoint string `kong:"default=otel:4318"`
}

func (c TraceConfig) Create(ctx context.Context) (func(), error) {
	exp, err := otlptracehttp.New(
		ctx,
		otlptracehttp.WithEndpoint(c.Endpoint),
		otlptracehttp.WithInsecure(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create trace exporter %w", err)
	}

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exp),
		trace.WithSampler(trace.AlwaysSample()),
	)

	otel.SetTracerProvider(tp)

	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			gcppropagator.CloudTraceFormatPropagator{},
			propagation.TraceContext{},
			propagation.Baggage{},
		))

	cleanup := func() {
		_ = tp.Shutdown(context.Background())
		_ = exp.Shutdown(context.Background())
	}
	return cleanup, nil
}

type MetricsConfig struct {
	Endpoint string `kong:"default=otel:4319"`
}

func (c MetricsConfig) Create(ctx context.Context) (func(), error) {
	exp, err := otlpmetrichttp.New(
		ctx,
		otlpmetrichttp.WithEndpoint(c.Endpoint),
		otlpmetrichttp.WithInsecure(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create metrics exporter %w", err)
	}

	pusher := controller.New(
		processor.NewFactory(
			simple.NewWithHistogramDistribution(),
			exp,
		),
		controller.WithExporter(exp),
		controller.WithCollectPeriod(time.Second),
	)

	if err := pusher.Start(context.Background()); err != nil {
		return nil, err
	}

	global.SetMeterProvider(pusher)

	cleanup := func() {
		if err := pusher.Stop(context.Background()); err != nil {
			otel.Handle(err)
		}

		if err := exp.Shutdown(context.Background()); err != nil {
			otel.Handle(err)
		}
	}

	return cleanup, nil
}

type Creator interface {
	Create(ctx context.Context) (func(), error)
}

func Create(ctx context.Context, creators ...Creator) (func(), error) {
	var cleanups []func()

	for _, c := range creators {
		cleanup, err := c.Create(ctx)
		if err != nil {
			return nil, err
		}

		cleanups = append(cleanups, cleanup)
	}

	cleanup := func() {
		for i := range cleanups {
			cleanups[i]()
		}
	}

	return cleanup, nil
}
