<?php

namespace AppServer\Logging;

use Monolog\Formatter\JsonFormatter;
use Monolog\LogRecord;

class Formatter extends JsonFormatter
{
    public function __construct()
    {
        parent::__construct(JsonFormatter::BATCH_MODE_NEWLINES);
    }

    public function format(LogRecord $record): string
    {
        $normalized = parent::normalize($record);
        $entries = [];
        $entries = array_merge($entries, $normalized['extra'] ?? []);
        $entries = array_merge($entries, $normalized['context'] ?? []);
        $entries['message'] = $record->message;
        $entries['timestamp'] = $record->datetime->format($this->getDateFormat());
        $entries['severity'] = $record->level->getName();

        return $this->toJson($entries, true) . ($this->appendNewline ? "\n" : '');
    }
}