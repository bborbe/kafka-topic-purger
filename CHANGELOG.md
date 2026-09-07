# Changelog

All notable changes to this project will be documented in this file.

Please choose versions by [Semantic Versioning](http://semver.org/).

* MAJOR version when you make incompatible API changes,
* MINOR version when you add functionality in a backwards compatible manner, and
* PATCH version when you make backwards compatible bug fixes.

## Unreleased

- chore: update Go to 1.27.1

## v0.2.3

- fix: `make build` refuses to stamp a version onto a tree that is not that version's tag (`check-version-tag`, escape hatch `ALLOW_UNTAGGED_BUILD=1`). `VERSION` defaults to the newest tag repo-wide, so an operator-run build from an untagged or older tree silently republishes under the newest tag. The guard compares `git describe --exact-match HEAD` against `$(VERSION)` and exits non-zero on mismatch.

## v0.2.2

- ci: add ci.yml running `make precommit` -- the `test` required status check had no workflow producing it, permanently blocking every PR

- chore: update github.com/bborbe/kafka to v1.25.11, github.com/bborbe/metrics to v0.6.1, github.com/bborbe/run to v1.10.2

## v0.2.1

- chore: update github.com/IBM/sarama to v1.60.2, github.com/bborbe/errors to v1.6.0, github.com/bborbe/http to v1.26.25, github.com/bborbe/log to v1.6.25, github.com/bborbe/metrics to v0.5.15, github.com/bborbe/run to v1.10.0, github.com/bborbe/sentry to v1.10.0, github.com/bborbe/service to v1.10.10, github.com/bborbe/time to v1.27.11, github.com/onsi/gomega to v1.43.0

## v0.2.0

- feat: opt into `autoMerge.trivial` for mechanically-trivial update PRs

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
