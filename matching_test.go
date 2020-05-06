package backoff

import (
	"errors"
	"testing"
)

func Test_ErrorMatching(t *testing.T) {
	customError = errors.New("custom error")
	differentError = errors.New("different error")

	tests = map[error]error{
		// sanity check
		customError:    customError,
		differentError: differentError,
		// actual test
		Permanent(customError):    customError,
		Permanent(differentError): differentError,
	}

	for err, expectedType := range tests {
		if !errors.Is(err, expectedType) {
			t.Fatalf("errors.Is: expected %v to match %T", err, expectedType)
		}
		if !errors.As(err, &expectedType) {
			t.Fatalf("errors.As: expected %v to match %T", err, expectedType)
		}
	}
}
