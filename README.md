# results
Package results provides generic types and functions for handling function results.

[![Go Reference](https://pkg.go.dev/badge/github.com/goaux/results.svg)](https://pkg.go.dev/github.com/goaux/results)
[![Go Report Card](https://goreportcard.com/badge/github.com/goaux/results)](https://goreportcard.com/report/github.com/goaux/results)

This package is designed to be simple, efficient, and easy to use while
adhering to Go's idiomatic practices.

The package aims to reduce boilerplate code associated with error checking and
fallback value assignment. It is particularly useful for functions that return
multiple values and an error.

## May

Maybe1, Maybe2, and Maybe3 types encapsulate one, two, or three values
respectively, along with an error. These types simplify error handling and
provide a convenient way to work with functions that may fail.

May1, May2 and May3 functions return a value of these types.

The number "1", "2" and "3" in the function name refer to the number of values
excluding an error they return.

### Example

    value := results.May1(someFunction()).Or(defaultValue)
    v1, v2 := results.May2(anotherFunction()).Or(default1, default2)

## Must

Must1, Must2 and Must3 functions panic if an error occurs, which is useful for
operations that should never fail during normal execution.

The number "1", "2" and "3" in the function name refer to the number of values they return.

### Example

    value := results.Must1(someFunction())
    v1, v2 := results.Must2(anotherFunction())

## Take and Drop

Take2, Drop2, Take3 and Drop3 functions can reduce the results of functions.

Both Take2 and Take3 return only the first value.

Both Drop2 and Drop3 discard the first value, return remaining value(s).

The numbers "2" and "3" in the function names refer to the number of arguments they accept.

### Example

    first := results.Take2(someFunction())
    second := results.Drop2(someFunction())

    firstOf3 := results.Take3(anotherFunction())
    the2, the3 := results.Drop3(anotherFunction())

## If1 and If2

If1 and If2 functions simulate conditional expression.

### Example

    // someFunction returns a value of any type and a boolean value.
    value := results.If1(someFunction()).Or(valueForFalse)
    // anotherFunction returns two values of any type and a boolean value.
    v1, v2 := results.If2(anotherFunction()).Or(fallback0, fallback1)

## ErrorAs

ErrorAs returns the result of calling errors.As(err, &target).

The coding pattern using errors.As is following:

    var pathError *fs.PathError
    if errors.As(err, &pathError) {

Using ErrorAs is a bit more concise:

    if pathError, ok := results.ErrorAs[*fs.PathError](err); ok {

Improvements:

- Just one line.
- Create a scope for the pathError variable.

## TypeAs

TypeAs returns the result of type assertions.
This is a function, so it is addressable.
