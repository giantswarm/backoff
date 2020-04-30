package backoff

import (
	"errors"
	"fmt"
	"testing"
)

var (
	customError    = fmt.Errorf("custom error")
	differentError = fmt.Errorf("different error")

	tests = map[error]error{
		// sanity check
		customError:    customError,
		differentError: differentError,
		// actual test
		Permanent(customError):    customError,
		Permanent(differentError): differentError,
	}
)

func Test_ErrorMatching(t *testing.T) {
	for err, expectedType := range tests {
		if !errors.Is(err, expectedType) {
			t.Fatalf("expected %v to match %T", err, expectedType)
		}
	}
}
