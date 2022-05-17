<?php

namespace AppServer\Logging;

use Monolog\LogRecord;
use Monolog\Processor\ProcessorInterface;
use OpenTelemetry\Context\Context;
use OpenTelemetry\SDK\Trace\Span;

class TraceProcessor implements ProcessorInterface
{
    public function __invoke(LogRecord $record): LogRecord
    {
        $ctx = Context::getCurrent();
        $span = Span::fromContext($ctx);

        $ctx = $span->getContext();

        $record->extra['traceId'] = $ctx->getTraceId();

        return $record;
    }
}
