package results_test

import (
	"fmt"
	"strconv"

	"github.com/goaux/results"
)

var _ results.Maybe = results.Maybe1[any]{}

func ExampleMaybe1() {
	v := results.May1(strconv.Atoi("999")).Or(42)
	fmt.Println(v)

	v = results.May1(strconv.Atoi("999")).OrPanic()
	fmt.Println(v)

	v = results.May1(strconv.Atoi("NaN")).Or(42)
	fmt.Println(v)
	// Output:
	// 999
	// 999
	// 42
}
