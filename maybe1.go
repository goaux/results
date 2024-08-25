package results

import "github.com/goaux/stacktrace"

// Maybe1 holds a value and an error
type Maybe1[T any] struct {
	V0  T
	Err error
}

// May1 creates a Maybe1 from a v0 and an error
func May1[T any](v0 T, err error) Maybe1[T] {
	return Maybe1[T]{
		V0:  v0,
		Err: err,
	}
}

// Or returns the default value if there's an error, otherwise returns the stored value
func (m Maybe1[T]) Or(defaultValue T) T {
	if m.Err != nil {
		return defaultValue
	}
	return m.V0
}

// OrPanic panics if there's an error, otherwise returns the stored value
func (m Maybe1[T]) OrPanic() T {
	if m.Err != nil {
		panic(stacktrace.Format(m.Err))
	}
	return m.V0
}

// Values returns the stored value.
func (m Maybe1[T]) Values() []any {
	return []any{m.V0}
}

// Error returns the error
func (m Maybe1[T]) Error() error {
	return m.Err
}

// HasError returns true if there's an error, false otherwise
func (m Maybe1[T]) HasError() bool {
	return m.Err != nil
}
