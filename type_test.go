package results_test

import (
	"fmt"

	"github.com/goaux/results"
)

func ExampleTypeAs() {
	asStr := results.TypeAs[string]
	asInt := results.TypeAs[int]
	fmt.Println(asStr("hello"))
	fmt.Println(asInt("hello"))
	// Output:
	// hello true
	// 0 false
}
