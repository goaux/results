package results_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/goaux/results"
)

func ExampleTrue1() {
	os.Setenv("EXAMPLE", "example")
	os.Unsetenv("EXAMPLE2")
	fmt.Println(results.If1(os.LookupEnv("EXAMPLE")).Or("fallback"))
	fmt.Println(results.If1(os.LookupEnv("EXAMPLE")).OrPanic("undefined"))
	fmt.Println(results.If1(os.LookupEnv("EXAMPLE2")).Or("fallback"))
	fmt.Println(results.If1(os.LookupEnv("EXAMPLE2")).OrFunc(func() string { return "fallback" }))
	// Output:
	// example
	// example
	// fallback
	// fallback
}

func TestTrue1(t *testing.T) {
	t.Run("Or", func(t *testing.T) {
		if v := results.If1("TRUE", true).Or("FALSE"); v != "TRUE" {
			t.Errorf("got %v, want TRUE", v)
		}
		if v := results.If1("TRUE", false).Or("FALSE"); v != "FALSE" {
			t.Errorf("got %v, want FALSE", v)
		}
		if v := results.If1(111, true).Or(222); v != 111 {
			t.Errorf("got %v, want 111", v)
		}
		if v := results.If1(111, false).Or(222); v != 222 {
			t.Errorf("got %v, want 222", v)
		}
	})

	t.Run("OrFunc", func(t *testing.T) {
		f0 := func() string { return "FALSE" }
		f1 := func() int { return 222 }
		if v := results.If1("TRUE", true).OrFunc(f0); v != "TRUE" {
			t.Errorf("got %v, want TRUE", v)
		}
		if v := results.If1("TRUE", false).OrFunc(f0); v != "FALSE" {
			t.Errorf("got %v, want FALSE", v)
		}
		if v := results.If1(111, true).OrFunc(f1); v != 111 {
			t.Errorf("got %v, want 111", v)
		}
		if v := results.If1(111, false).OrFunc(f1); v != 222 {
			t.Errorf("got %v, want 222", v)
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
			_ = results.If1("", true).OrPanic("message")
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
			_ = results.If1("", false).OrPanic(want)
			t.Errorf("never")
		})
	})

	t.Run("Values", func(t *testing.T) {
		t.Run("true", func(t *testing.T) {
			got := results.If1(99, true).Values()
			want := []any{99}
			if !slicesEqual(got, want) {
				t.Errorf("must be equal, got %v, want %v", got, want)
			}
		})

		t.Run("false", func(t *testing.T) {
			got := results.If1(99, false).Values()
			want := []any{99}
			if !slicesEqual(got, want) {
				t.Errorf("must be equal, got %v, want %v", got, want)
			}
		})
	})

	t.Run("IsOK", func(t *testing.T) {
		t.Run("true", func(t *testing.T) {
			got := results.If1(99, true).IsOK()
			if !got {
				t.Errorf("must be true")
			}
		})

		t.Run("false", func(t *testing.T) {
			got := results.If1(99, false).IsOK()
			if got {
				t.Errorf("must be false")
			}
		})
	})
}
