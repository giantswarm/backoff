package backoff

import (
	"time"

	"github.com/cenkalti/backoff"
)

func NewConstantBackoff(maxWait, maxInterval time.Duration) backoff.BackOff {
	b := withMaxElapsedTime(backoff.NewConstantBackOff(maxInterval), maxWait)

	b.Reset()

	return b
}

func withMaxElapsedTime(b backoff.BackOff, d time.Duration) *backOffMaxElapsedTime {
	return &backOffMaxElapsedTime{
		underlying: b,
		maxElapsed: d,
		start:      time.Time{},
	}
}

type backOffMaxElapsedTime struct {
	underlying backoff.BackOff
	maxElapsed time.Duration
	start      time.Time
}

func (b *backOffMaxElapsedTime) NextBackOff() time.Duration {
	if b.start.IsZero() {
		b.start = time.Now()
	}

	if time.Now().After(b.start.Add(b.maxElapsed)) {
		return backoff.Stop
	}

	return b.underlying.NextBackOff()
}

func (b *backOffMaxElapsedTime) Reset() {
	b.start = time.Time{}
	b.underlying.Reset()
}