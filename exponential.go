package backoff

import (
	"time"

	"github.com/cenkalti/backoff/v6"
)

func NewExponential(maxWait, maxInterval time.Duration) BackOff {
	e := backoff.NewExponentialBackOff()
	e.MaxInterval = maxInterval

	b := withMaxElapsedTime(e, maxWait)

	b.Reset()

	return b
}
