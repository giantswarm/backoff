package backoff

import (
	"errors"
	"testing"
)

type simpleTestError struct{}

func (c *simpleTestError) Error() string {
	return "simpleTestError"
}

type wrapTestError struct {
	err error
}

func (c *wrapTestError) Error() string {
	return c.err.Error()
}

func (c *wrapTestError) Unwrap() error {
	return c.err
}

func Test_ErrorMatching(t *testing.T) {
	customError := errors.New("custom error")
	differentError := errors.New("different error")
	customTypeError := &simpleTestError{}
	wrappedError := &wrapTestError{customTypeError}

	tests := map[error]error{
		// sanity check
		customError:    customError,
		differentError: differentError,
		// actual test
		Permanent(customError):     customError,
		Permanent(differentError):  differentError,
		Permanent(customTypeError): customTypeError,
		Permanent(wrappedError):    wrappedError,
		Permanent(wrappedError):    customTypeError,
	}

	var err, expectedTarget error
	for err, expectedTarget = range tests {
		if !errors.Is(err, expectedTarget) {
			t.Fatalf("errors.Is: expected %v to match %T", err, expectedTarget)
		}
	}
}
