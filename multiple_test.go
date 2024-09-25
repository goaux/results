package results_test

import (
	"fmt"
	"net"
	"testing"

	"github.com/goaux/results"
)

func ExampleGet1of2() {
	fmt.Println(results.Get1of2(
		results.May2(net.SplitHostPort("example.com:80")).Or("localhost", "80"),
	))
	// Output:
	// example.com
}

func ExampleGet2of2() {
	fmt.Println(results.Get2of2(
		results.May2(net.SplitHostPort("example.com:80")).Or("localhost", "80"),
	))
	// Output:
	// 80
}

func ExampleGet1of3() {
	example := func() (x, y, z int) { return 11, 22, 33 }
	fmt.Println(results.Get1of3(example()))
	// Output:
	// 11
}

func ExampleGet23of3() {
	example := func() (x, y, z int) { return 11, 22, 33 }
	fmt.Println(results.Get23of3(example()))
	// Output:
	// 22 33
}

func TestGet1of2(t *testing.T) {
	result := results.Get1of2(1, "two")
	if result != 1 {
		t.Errorf("Get1of2(1, \"two\") = %v; want 1", result)
	}
}

func TestGet2of2(t *testing.T) {
	result := results.Get2of2(1, "two")
	if result != "two" {
		t.Errorf("Get2of2(1, \"two\") = %v; want \"two\"", result)
	}
}

func TestGet1of3(t *testing.T) {
	result := results.Get1of3(1, "two", 3.0)
	if result != 1 {
		t.Errorf("Get1of3(1, \"two\", 3.0) = %v; want 1", result)
	}
}

func TestGet2of3(t *testing.T) {
	result := results.Get2of3(1, "two", 3.0)
	if result != "two" {
		t.Errorf("Get2of3(1, \"two\", 3.0) = %v; want \"two\"", result)
	}
}

func TestGet3of3(t *testing.T) {
	result := results.Get3of3(1, "two", 3.0)
	if result != 3.0 {
		t.Errorf("Get3of3(1, \"two\", 3.0) = %v; want 3.0", result)
	}
}

func TestGet12of3(t *testing.T) {
	a, b := results.Get12of3(1, "two", 3.0)
	if a != 1 || b != "two" {
		t.Errorf("Get12of3(1, \"two\", 3.0) = %v, %v; want 1, \"two\"", a, b)
	}
}

func TestGet23of3(t *testing.T) {
	b, c := results.Get23of3(1, "two", 3.0)
	if b != "two" || c != 3.0 {
		t.Errorf("Get23of3(1, \"two\", 3.0) = %v, %v; want \"two\", 3.0", b, c)
	}
}
