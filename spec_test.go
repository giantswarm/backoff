package backoff

import (
	"testing"

	"github.com/cenkalti/backoff"
)

// Test_BackOff tests if this library and underlying implementation interfaces
// are compatible.
func Test_BackOff(t *testing.T) {
	var custom BackOff
	var underlying backoff.BackOff

	// Create an anonymus function that takes the interface we want to compare
	// against as argument and provide the other implementation in order to verify
	// it implements the actual interface. Doing that for both interfaces verifies
	// both interfaces are equal.
	func(v BackOff) {}(underlying)
	func(v backoff.BackOff) {}(custom)
}
