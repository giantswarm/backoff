package backoff

import (
	"github.com/cenkalti/backoff/v5"
)

func Permanent(err error) error {
	return backoff.Permanent(err)
}
