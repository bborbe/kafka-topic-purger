# kafka-topic-purger

HTTP service that purges all messages from a Kafka topic. Broker-agnostic — works with any Kafka cluster.

## Run

```
make run
```

Purge a topic:

```
curl -X POST http://localhost:20032/purge/<topic>
```

## Endpoints

- `POST /purge/{topic}` — purge all messages from the given topic
- `GET  /healthz` / `GET /readiness` — health checks
- `GET  /metrics` — Prometheus metrics
- `GET  /setloglevel/{level}` — adjust log level at runtime

## Flags

| Flag | Env | Description |
|---|---|---|
| `-kafka-brokers` | `KAFKA_BROKERS` | comma-separated Kafka brokers |
| `-listen` | `LISTEN` | listen address (e.g. `:20032`) |
| `-dry-run` | `DRY_RUN` | log what would be purged without deleting |
| `-sentry-dsn` | `SENTRY_DSN` | optional Sentry DSN |

## Build

`make buca` builds and publishes `docker.io/bborbe/kafka-topic-purger:vX.Y.Z` (git-tag semver).

## License

This project is licensed under the BSD-style license. See the [LICENSE](LICENSE) file for details.
