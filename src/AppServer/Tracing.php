<?php

namespace AppServer;

use GuzzleHttp\Client;
use GuzzleHttp\Psr7\HttpFactory;
use OpenTelemetry\SDK\Trace\SpanProcessor\BatchSpanProcessor;
use OpenTelemetry\SDK\Trace\SpanProcessor\SimpleSpanProcessor;
use OpenTelemetry\SDK\Trace\TracerProvider;
use OpenTelemetry\Contrib\OtlpHttp\Exporter;

class Tracing
{
    public static function initialize(): void
    {
        $factory = new HttpFactory();

        $exporter = new Exporter(
            new Client(),
            $factory,
            $factory
        );

        $processor = new BatchSpanProcessor(
            //new SimpleSpanProcessor($exporter),
            $exporter,
            //null,
            //32,
           // 1000,
        );

        $tracerProvider = new TracerProvider($processor);

        $tracer = $tracerProvider->getTracer();

        TracerProvider::setDefaultTracer($tracer);
    }

    public static function startSpan(string $name): \OpenTelemetry\API\Trace\SpanInterface
    {
        return TracerProvider::getDefaultTracer()->spanBuilder($name)->startSpan();
    }
}