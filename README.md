[![CircleCI](https://dl.circleci.com/status-badge/img/gh/giantswarm/backoff/tree/main.svg?style=svg)](https://dl.circleci.com/status-badge/redirect/gh/giantswarm/backoff/tree/main)

# backoff

Backoff is a library abstracting retry functionality using
[github.com/cenkalti/backoff](https://pkg.go.dev/github.com/cenkalti/backoff). We use this library to have a
unified interface and the opportunity to wrap and extend backoff implementations
as we need them.
