# Changelog

All notable changes to this project will be documented in this file.

Please choose versions by [Semantic Versioning](http://semver.org/).

* MAJOR version when you make incompatible API changes,
* MINOR version when you add functionality in a backwards compatible manner, and
* PATCH version when you make backwards compatible bug fixes.

## v0.1.4

- chore: Bump errcheck to v1.20.0 and golangci-lint to v2.13.1 for Go 1.27 support
## v0.1.3

- update Go to 1.26.6 and update dependencies

## v0.1.2

- update Go to 1.26.6 and update dependencies, fixing GO-2026-6179 and GO-2026-6180 in golang.org/x/mod

## v0.1.1

- docs: add a License section to the README

## v0.1.0

- Initial release — extracted from `bborbe/trading` (`strimzi/topic-purger`) as a standalone public repo
- HTTP service purging all messages from a Kafka topic (`POST /purge/{topic}`), broker-agnostic
- Decoupled from `trading/lib`: build-info via public `github.com/bborbe/metrics`, log-level + sync-producer via public `github.com/bborbe/log` / `github.com/bborbe/kafka`
- Publish-only build → `docker.io/bborbe/kafka-topic-purger:vX.Y.Z`
