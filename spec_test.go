package backoff

import (
	"testing"

	"github.com/cenkalti/backoff/v3"
)

// Test_BackOff tests if this library and underlying implementation
// interfaces are compatible.
func Test_BackOff(t *testing.T) {
	var custom BackOff
	var underlying backoff.BackOff
	custom = underlying
	underlying = custom
}
