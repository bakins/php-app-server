<?php
require __DIR__ . '/../php-vendor/autoload.php';

use AppServer\Logging;
use AppServer\Logging\MonologFormatter;
use AppServer\OpenTelemetry\Monolog\Handler;
use AppServer\Tracing;
use DataDog\DogStatsd;
use GuzzleHttp\Psr7\ServerRequest;
use Laminas\HttpHandlerRunner\Emitter\SapiEmitter;
use OpenTelemetry\SDK\Common\Log\LoggerHolder;
use Twirp\Router;
use Twitch\Twirp\Example\HaberdasherServer;
use Twitch\Twirp\Example\Hat;
use Twitch\Twirp\Example\Size;

Logging::initialize();
Tracing::initialize();

final class Haberdasher implements \Twitch\Twirp\Example\Haberdasher
{
    private DogStatsd $statsd;

    public function __construct()
    {
        $this->statsd = new DogStatsd(['host' => 'statsd', 'port' => 8125, 'global_tags' => ['service' => 'Haberdasher']]);
    }

    public function MakeHat(array $ctx, Size $size): Hat
    {
        $span = Tracing::startSpan('MakeHat');

        $hat = new Hat();
        $hat->setSize($size->getInches());
        $hat->setColor('golden');
        $hat->setName('crown');

        LoggerHolder::get()->warning('hello world', ['this' => 'that', 'foo' => ["a", "b"]]);

        $span->end();

        $this->statsd->increment('hats_made', 1.0, ['color' => $hat->getName()], 1);

        return $hat;
    }
}

$span = Tracing::startSpan('root');
$span->activate();

$router = new Router();

$haberdasher = new HaberdasherServer(new Haberdasher());
#, new TraceHooks());

$router->registerHandler($haberdasher->getPathPrefix(), $haberdasher);

$request = ServerRequest::fromGlobals();

$response = $router->handle($request);

(new SapiEmitter())->emit($response);

$span->end();
