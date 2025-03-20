package results_test

import (
	"errors"
	"fmt"
	"io/fs"
	"strconv"
	"testing"

	"github.com/goaux/results"
)

func ExampleErrorAs() {
	err := fmt.Errorf("wrap %w", &fs.PathError{Op: "example"})
	if pathError, ok := results.ErrorAs[*fs.PathError](err); ok {
		fmt.Println(pathError.Op)
	}
	// Output:
	// example
}

func TestErrorAs(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		err := fmt.Errorf("wrap %w", &fs.PathError{Op: "example"})
		pathError, ok := results.ErrorAs[*fs.PathError](err)
		if !ok {
			t.Error("should be true")
		}
		if pathError == nil {
			t.Error("should not be nil")
		}
		if pathError.Op != "example" {
			t.Error("should be equal")
		}
	})
	t.Run("false", func(t *testing.T) {
		err := fmt.Errorf("wrap %w", fs.ErrClosed)
		pathError, ok := results.ErrorAs[*fs.PathError](err)
		if ok {
			t.Error("should be false")
		}
		if pathError != nil {
			t.Error("should be nil")
		}
	})
}

func TestErrorIs(t *testing.T) {
	Err1 := errors.New("Err1")
	Err2 := errors.New("Err2")
	Err3 := errors.New("Err3")

	WErr1 := fmt.Errorf("WErr1 %w", Err1)
	WErr2 := fmt.Errorf("WErr2 %w", Err2)
	WErr3 := fmt.Errorf("WErr3 %w", Err3)

	WErr12 := fmt.Errorf("WErr12 %w %w", Err1, Err2)

	tests := []struct {
		err    error
		target []error
		expect bool
	}{
		{Err1, []error{}, false},
		{Err1, []error{Err1}, true},
		{Err1, []error{Err2}, false},
		{Err1, []error{Err1, Err2}, true},
		{Err2, []error{Err1, Err2}, true},
		{Err3, []error{Err1, Err2}, false},

		{WErr1, []error{}, false},
		{WErr1, []error{Err1}, true},
		{WErr1, []error{Err2}, false},
		{WErr1, []error{Err1, Err2}, true},
		{WErr2, []error{Err1, Err2}, true},
		{WErr3, []error{Err1, Err2}, false},

		{WErr12, []error{}, false},
		{WErr12, []error{Err1}, true},
		{WErr12, []error{Err2}, true},
		{WErr12, []error{Err1, Err2}, true},
		{WErr12, []error{Err3}, false},
	}
	for i, tc := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := results.ErrorIs(tc.err, tc.target...)
			if actual != tc.expect {
				t.Errorf("actual=%v expect=%v", actual, tc.expect)
			}
		})
	}
}
