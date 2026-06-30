package backoff

import (
	"github.com/cenkalti/backoff/v7"
)

func Permanent(err error) error {
	return backoff.Permanent(err)
}
