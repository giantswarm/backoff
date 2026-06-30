package backoff

import (
	"context"

	"github.com/cenkalti/backoff/v6"
	"github.com/giantswarm/microerror"
)

// Retry retries the operation o until it does not return error or BackOff
// stops. See https://godoc.org/github.com/cenkalti/backoff#Retry for details.
func Retry(o Operation, b BackOff) error {
	_, err := backoff.Retry(
		context.Background(),
		func() (struct{}, error) { return struct{}{}, o() },
		backoff.WithBackOff(b),
		backoff.WithMaxElapsedTime(0),
	)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}

// RetryNotify does what Retry do with notification between each try.
func RetryNotify(o Operation, b BackOff, n Notify) error {
	_, err := backoff.Retry(
		context.Background(),
		func() (struct{}, error) { return struct{}{}, o() },
		backoff.WithBackOff(b),
		backoff.WithNotify(backoff.Notify(n)),
		backoff.WithMaxElapsedTime(0),
	)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}
