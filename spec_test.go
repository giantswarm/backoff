package backoff

import (
	"testing"

	"github.com/cenkalti/backoff"
)

// Test_Interface tests if this library and underlying implementation
// interfaces are compatible.
func Test_Interface(t *testing.T) {
	var custom BackOff
	var underlying backoff.BackOff
	custom = underlying
	underlying = custom
}
