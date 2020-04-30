package backoff

import (
	"errors"
	"fmt"
	"testing"
)

var (
	customError    = fmt.Errorf("custom error")
	differentError = fmt.Errorf("different error")
)

func Test_ErrorMatching(t *testing.T) {
	// sanity check
	err1 := customError
	if !errors.Is(err1, customError) {
		t.Fatalf("expected %v to match %T", err1, customError)
	}

	err2 := differentError
	if !errors.Is(err2, differentError) {
		t.Fatalf("expected %v to match %T", err2, differentError)
	}

	// actual test
	permanentErr1 := Permanent(err1)
	if !errors.Is(permanentErr1, customError) {
		t.Fatalf("expected %v to match %T", permanentErr1, customError)
	}

	permanentErr2 := Permanent(err2)
	if !errors.Is(permanentErr2, differentError) {
		t.Fatalf("expected %v to match %T", permanentErr2, differentError)
	}
}
