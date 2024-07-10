package results_test

import (
	"fmt"
	"net"

	"github.com/goaux/results"
)

var _ results.Maybe = results.Maybe2[any, any]{}

func ExampleMaybe2() {
	host, port := results.May2(net.SplitHostPort("example.com:80")).Or("localhost", "80")
	fmt.Println(host, port)

	host, port = results.May2(net.SplitHostPort("example.com:80")).OrPanic()
	fmt.Println(host, port)

	host, port = results.May2(net.SplitHostPort("error")).Or("localhost", "80")
	fmt.Println(host, port)
	// Output:
	// example.com 80
	// example.com 80
	// localhost 80
}
