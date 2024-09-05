package results_test

import (
	"fmt"
	"net"
	"testing"

	"github.com/goaux/results"
)

func ExampleTake2() {
	fmt.Println(results.Take2(
		results.May2(net.SplitHostPort("example.com:80")).Or("localhost", "80"),
	))
	// Output:
	// example.com
}

func ExampleDrop2() {
	fmt.Println(results.Drop2(
		results.May2(net.SplitHostPort("example.com:80")).Or("localhost", "80"),
	))
	// Output:
	// 80
}

func ExampleTake3() {
	example := func() (x, y, z int) { return 11, 22, 33 }
	fmt.Println(results.Take3(example()))
	// Output:
	// 11
}

func ExampleDrop3() {
	example := func() (x, y, z int) { return 11, 22, 33 }
	fmt.Println(results.Drop3(example()))
	// Output:
	// 22 33
}

func TestTake2(t *testing.T) {
	if results.Take2(func() (int, int) { return 11, 22 }()) != 11 {
		t.Error()
	}
}

func TestDrop2(t *testing.T) {
	if results.Drop2(func() (int, int) { return 11, 22 }()) != 22 {
		t.Error()
	}
}

func TestTake3(t *testing.T) {
	if results.Take3(func() (a, b, c int) { return 11, 22, 33 }()) != 11 {
		t.Error()
	}
}

func TestDrop3(t *testing.T) {
	if b, c := results.Drop3(func() (a, b, c int) { return 11, 22, 33 }()); b != 22 || c != 33 {
		t.Error()
	}
}
