# results
Package results provides generic types for handling function results with potential errors.

[![Go Reference](https://pkg.go.dev/badge/github.com/goaux/results.svg)](https://pkg.go.dev/github.com/goaux/results)

It offers Maybe1, Maybe2, and Maybe3 types that encapsulate one, two, or three values
respectively, along with an error. These types simplify error handling and provide
a convenient way to work with functions that may fail.

The package aims to reduce boilerplate code associated with error checking and
fallback value assignment. It is particularly useful for functions that return
multiple values and an error.

Usage:

    value := results.May1(someFunction()).Or(defaultValue)

    v1, v2 := results.May2(anotherFunction()).Or(default1, default2)

This package is designed to be simple, efficient, and easy to use while
adhering to Go's idiomatic practices.
