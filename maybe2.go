package results

import "github.com/goaux/stacktrace"

// Maybe2 holds two values and an error
type Maybe2[T0, T1 any] struct {
	V0  T0
	V1  T1
	Err error
}

// May2 creates a Maybe2 from two values and an error
func May2[T0, T1 any](v0 T0, v1 T1, err error) Maybe2[T0, T1] {
	return Maybe2[T0, T1]{
		V0:  v0,
		V1:  v1,
		Err: err,
	}
}

// Or returns the default values if there's an error, otherwise returns the stored values
func (m Maybe2[T0, T1]) Or(d0 T0, d1 T1) (T0, T1) {
	if m.Err != nil {
		return d0, d1
	}
	return m.V0, m.V1
}

// OrPanic panics if there's an error, otherwise returns the stored values
func (m Maybe2[T0, T1]) OrPanic() (T0, T1) {
	if m.Err != nil {
		panic(stacktrace.Format(m.Err))
	}
	return m.V0, m.V1
}

// Values returns the stored values.
func (m Maybe2[T0, T1]) Values() []any {
	return []any{m.V0, m.V1}
}

// Error returns the error
func (m Maybe2[T0, T1]) Error() error {
	return m.Err
}

// HasError returns true if there's an error, false otherwise
func (m Maybe2[T0, T1]) HasError() bool {
	return m.Err != nil
}
