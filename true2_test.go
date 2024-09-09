package results_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/goaux/results"
)

func ExampleTrue2() {
	fmt.Println(results.If2("example", 42, true).Or("fallback", 99))
	fmt.Println(results.If2("example", 42, true).OrPanic("undefined"))
	fmt.Println(results.If2("example", 42, false).Or("fallback", 99))
	fmt.Println(results.If2("example", 42, false).OrFunc(func() (string, int) { return "fallback", 99 }))
	// Output:
	// example 42
	// example 42
	// fallback 99
	// fallback 99
}

func TestTrue2(t *testing.T) {
	t.Run("Or", func(t *testing.T) {
		if a, b := results.If2(1, 2, true).Or(3, 4); a != 1 || b != 2 {
			t.Errorf("got %v %v, want 1 2", a, b)
		}
		if a, b := results.If2(1, 2, false).Or(3, 4); a != 3 || b != 4 {
			t.Errorf("got %v %v, want 3 4", a, b)
		}
		if a, b := results.If2("1", 2, true).Or("3", 4); a != "1" || b != 2 {
			t.Errorf("got %v %v, want 1 2", a, b)
		}
		if a, b := results.If2("1", 2, false).Or("3", 4); a != "3" || b != 4 {
			t.Errorf("got %v %v, want 3 4", a, b)
		}
	})

	t.Run("OrFunc", func(t *testing.T) {
		f0 := func() (string, int) { return "FALSE", 999 }
		f1 := func() (int, int) { return 888, 999 }
		if a, b := results.If2("TRUE", 42, true).OrFunc(f0); a != "TRUE" || b != 42 {
			t.Errorf("got %v %v, want TRUE 42", a, b)
		}
		if a, b := results.If2("TRUE", 42, false).OrFunc(f0); a != "FALSE" || b != 999 {
			t.Errorf("got %v %v, want FALSE 999", a, b)
		}
		if a, b := results.If2(12, 42, true).OrFunc(f1); a != 12 || b != 42 {
			t.Errorf("got %v %v, want 12 42", a, b)
		}
		if a, b := results.If2(12, 42, false).OrFunc(f1); a != 888 || b != 999 {
			t.Errorf("got %v %v, want 888 999", a, b)
		}
	})

	t.Run("OrPanic", func(t *testing.T) {
		t.Run("true", func(t *testing.T) {
			defer func() {
				v := recover()
				if v != nil {
					t.Errorf("got %v, should be nil", v)
				}
			}()
			_, _ = results.If2("", 42, true).OrPanic("message")
		})

		t.Run("false", func(t *testing.T) {
			want := "message"
			defer func() {
				v := recover()
				if v == nil {
					t.Error("should not be nil")
				}
				msg, ok := v.(string)
				if !ok {
					t.Error("should be string")
				}
				if msg != want {
					t.Errorf("got:%q want:%q", msg, want)
				}
			}()
			_, _ = results.If2("", 42, false).OrPanic(want)
			t.Errorf("never")
		})
	})

	t.Run("Values", func(t *testing.T) {
		t.Run("true", func(t *testing.T) {
			got := results.If2(99, 4.2, true).Values()
			want := []any{99, 4.2}
			if !slices.Equal(got, want) {
				t.Errorf("must be equal, got %v, want %v", got, want)
			}
		})

		t.Run("false", func(t *testing.T) {
			got := results.If2(99, 4.2, false).Values()
			want := []any{99, 4.2}
			if !slices.Equal(got, want) {
				t.Errorf("must be equal, got %v, want %v", got, want)
			}
		})
	})

	t.Run("IsOK", func(t *testing.T) {
		t.Run("true", func(t *testing.T) {
			got := results.If2(99, "", true).IsOK()
			if !got {
				t.Errorf("must be true")
			}
		})

		t.Run("false", func(t *testing.T) {
			got := results.If2("", 99, false).IsOK()
			if got {
				t.Errorf("must be false")
			}
		})
	})
}
