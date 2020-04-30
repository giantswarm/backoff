package backoff

import (
	"github.com/cenkalti/backoff/v4"
)

func Permanent(err error) error {
	return backoff.Permanent(err)
}
