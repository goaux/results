package results_test

import (
	"fmt"
	"testing"

	"github.com/goaux/results"
)

func ExampleIf() {
	fmt.Println(results.If(true, 42, 0))
	// Output:
	// 42
}

func ExampleIfFunc() {
	fmt.Println(
		results.IfFunc(
			true,
			func() string { return "42" },
			func() string { return "" },
		),
	)
	// Output:
	// 42
}

func TestIf(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		got := results.If(true, 1, 2)
		if got != 1 {
			t.Error("must be 1")
		}
	})
	t.Run("false", func(t *testing.T) {
		got := results.If(false, 1, 2)
		if got != 2 {
			t.Error("must be 2")
		}
	})
}

func TestIfFunc(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		got := results.IfFunc(true, func() int { return 1 }, func() int { return 2 })
		if got != 1 {
			t.Error("must be 1")
		}
	})
	t.Run("false", func(t *testing.T) {
		got := results.IfFunc(false, func() int { return 1 }, func() int { return 2 })
		if got != 2 {
			t.Error("must be 2")
		}
	})
}
