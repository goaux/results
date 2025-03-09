package results

import "github.com/goaux/stacktrace/v2"

// Must panics if the error is non-nil.
// It's useful for operations that should never fail during normal execution.
func Must(err error) {
	if err != nil {
		panic(stacktrace.Format(err))
	}
}

// Must1 panics if the error is non-nil, otherwise returns the value.
// It's useful for operations that should never fail during normal execution.
func Must1[T any](v T, err error) T {
	if err != nil {
		panic(stacktrace.Format(err))
	}
	return v
}

// Must2 panics if the error is non-nil, otherwise returns the two values.
// It's useful for operations that should never fail during normal execution.
func Must2[T0, T1 any](v0 T0, v1 T1, err error) (T0, T1) {
	if err != nil {
		panic(stacktrace.Format(err))
	}
	return v0, v1
}

// Must3 panics if the error is non-nil, otherwise returns the three values.
// It's useful for operations that should never fail during normal execution.
func Must3[T0, T1, T2 any](v0 T0, v1 T1, v2 T2, err error) (T0, T1, T2) {
	if err != nil {
		panic(stacktrace.Format(err))
	}
	return v0, v1, v2
}
