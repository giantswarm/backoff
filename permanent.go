package backoff

import (
	"github.com/cenkalti/backoff/v3"
)

func Permanent(err error) error {
	return backoff.Permanent(err)
}
