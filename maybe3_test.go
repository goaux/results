package results_test

import (
	"fmt"

	"github.com/goaux/results"
)

var _ results.Maybe = results.Maybe3[any, any, any]{}

func ExampleMaybe3() {
	example := func(string) (x, y, z int, err error) { return 8, 7, 6, nil }
	x, y, z := results.May3(example("8,7,6")).Or(1, 1, 1)
	fmt.Println(x, y, z)

	x, y, z = results.May3(example("8,7,6")).OrPanic()
	fmt.Println(x, y, z)

	example = func(string) (x, y, z int, err error) { return 0, 0, 0, fmt.Errorf("invalid") }
	x, y, z = results.May3(example("error")).Or(1, 1, 1)
	fmt.Println(x, y, z)
	// Output:
	// 8 7 6
	// 8 7 6
	// 1 1 1
}
