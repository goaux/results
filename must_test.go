package results_test

import (
	"errors"
	"testing"

	"github.com/goaux/results"
)

func TestMust(t *testing.T) {
	// Test successful case
	results.Must(func() error { return nil }())

	// Test panic case
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	results.Must(errors.New("test error"))
}

func TestMust1(t *testing.T) {
	// Test successful case
	v := results.Must1(42, nil)
	if v != 42 {
		t.Errorf("Expected 42, got %v", v)
	}

	// Test panic case
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	results.Must1(0, errors.New("test error"))
}

func TestMust2(t *testing.T) {
	// Test successful case
	v1, v2 := results.Must2("hello", 42, nil)
	if v1 != "hello" || v2 != 42 {
		t.Errorf("Expected 'hello' and 42, got %v and %v", v1, v2)
	}

	// Test panic case
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	results.Must2("", 0, errors.New("test error"))
}

func TestMust3(t *testing.T) {
	// Test successful case
	v1, v2, v3 := results.Must3("hello", 42, true, nil)
	if v1 != "hello" || v2 != 42 || v3 != true {
		t.Errorf("Expected 'hello', 42, and true, got %v, %v, and %v", v1, v2, v3)
	}

	// Test panic case
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	results.Must3("", 0, false, errors.New("test error"))
}
