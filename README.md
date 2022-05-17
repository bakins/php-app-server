# php-app-server

Example of running a PHP app.

## Components

See [docker-compose.yaml](./docker-compose.yaml) for the containers.

* `frontend` - HTTP server that forwards application requests to php using fastcgi.
* `phpfpm` -  Handles running the php application using [php-fpm](https://www.php.net/manual/en/install.fpm.php). The php application is in [example](./example)
* `statsd` - [stastd_exporter](https://github.com/prometheus/statsd_exporter) - the php app uses statsd to record metrics.
* `otel` - [opentelemetry-collector](https://opentelemetry.io/docs/collector/) - accepts traces from the php and Go processes. Accepts metrics from Go. Scrapes the statsd_exporter prometheus metrics.

Everything logs to stdout/stderr.

## TODO

* simple sqlite interaction
* sqlite replication using [litestream](https://litestream.io/
)
* HTTP/1 -> 2 proxy example.
* general code cleanup. This was written hastily to play with a few ideas.