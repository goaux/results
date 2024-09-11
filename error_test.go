package results_test

import (
	"fmt"
	"io/fs"
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
