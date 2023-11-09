# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Changed

- Upgrade to Go 1.20
- Upgrade github.com/giantswarm/microerror v0.4.0 to v0.4.1
- Upgrade github.com/giantswarm/micrologger v1.0.0 to v1.1.1
- Upgrade github.com/go-kit/log v0.2.0 to v0.2.1
- Upgrade github.com/go-logfmt/logfmt v0.5.1 to v0.6.0
- Upgrade github.com/go-logr/logr v1.2.2 to v1.3.0

## [1.0.0] - 2021-12-16

### Changed

- Upgrade to Go 1.17
- Upgrade github.com/giantswarm/microerror v0.2.0 to v0.4.0
- Upgrade github.com/giantswarm/micrologger v0.3.1 to v0.6.0
- Upgrade github.com/cenkalti/backoff v2.2.1 to v4.0.2
- Upgrade github.com/giantswarm/architect-orb v0.8.3 to v4.9.0

### Fixed

- Permanent backoff works with Golang error matching functions.

## [0.2.0] 2020-03-20

### Changed

- Switch from dep to Go modules.
- Use architect orb.

## [0.1.0] 2020-03-17

### Added

- First release.

[Unreleased]: https://github.com/giantswarm/backoff/compare/v1.0.0...HEAD
[1.0.0]: https://github.com/giantswarm/backoff/compare/v0.2.0...v1.0.0
[0.2.0]: https://github.com/giantswarm/backoff/releases/tag/v0.1.0...v0.2.0
[0.1.0]: https://github.com/giantswarm/backoff/releases/tag/v0.1.0
