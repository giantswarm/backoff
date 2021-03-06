package backoff

import (
	"time"

	"github.com/cenkalti/backoff/v4"
)

func NewMaxRetries(maxRetries uint64, maxInterval time.Duration) BackOff {
	b := withMaxRetries(backoff.NewConstantBackOff(maxInterval), maxRetries)

	b.Reset()

	return b
}

func withMaxRetries(b BackOff, m uint64) *backOffMaxRetries {
	return &backOffMaxRetries{
		maxRetries: m,
		retryCount: 0,
		underlying: b,
	}
}

type backOffMaxRetries struct {
	maxRetries uint64
	retryCount uint64
	underlying BackOff
}

func (b *backOffMaxRetries) NextBackOff() time.Duration {
	if b.retryCount+1 >= b.maxRetries {
		return backoff.Stop
	}
	b.retryCount++

	return b.underlying.NextBackOff()
}

func (b *backOffMaxRetries) Reset() {
	b.retryCount = 0
	b.underlying.Reset()
}
