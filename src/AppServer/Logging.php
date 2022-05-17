<?php

namespace AppServer;

use AppServer\Logging\Formatter;
use AppServer\Logging\TraceProcessor;
use Monolog\Handler\StreamHandler;
use Monolog\Level;
use Monolog\Logger;
use OpenTelemetry\SDK\Common\Log\LoggerHolder;

class Logging
{
    public static function initialize(string $name = 'application', Level $level = Level::Info): void
    {
        $logger = new Logger($name);

        register_shutdown_function([$logger, 'close']);

        $formatter = new Formatter();
        $formatter->setDateFormat(\DateTime::RFC3339);

        $handler = new StreamHandler('php://stdout', Level::Info);
        $handler->setFormatter($formatter);

        $logger->pushHandler($handler);

        $logger->pushProcessor(new TraceProcessor());

        LoggerHolder::set($logger);
    }
}
