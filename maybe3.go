package results

import "github.com/goaux/stacktrace"

// Maybe3 holds three values and an error
type Maybe3[T0, T1, T2 any] struct {
	V0  T0
	V1  T1
	V2  T2
	Err error
}

// May3 creates a Maybe3 from three values and an error
func May3[T0, T1, T2 any](v0 T0, v1 T1, v2 T2, err error) Maybe3[T0, T1, T2] {
	return Maybe3[T0, T1, T2]{
		V0:  v0,
		V1:  v1,
		V2:  v2,
		Err: err,
	}
}

// Or returns the default values if there's an error, otherwise returns the stored values
func (m Maybe3[T0, T1, T2]) Or(d0 T0, d1 T1, d2 T2) (T0, T1, T2) {
	if m.Err != nil {
		return d0, d1, d2
	}
	return m.V0, m.V1, m.V2
}

// OrPanic panics if there's an error, otherwise returns the stored values
func (m Maybe3[T0, T1, T2]) OrPanic() (T0, T1, T2) {
	if m.Err != nil {
		panic(stacktrace.Format(m.Err))
	}
	return m.V0, m.V1, m.V2
}

// Values returns the stored values.
func (m Maybe3[T0, T1, T2]) Values() []any {
	return []any{m.V0, m.V1, m.V2}
}

// Error returns the error
func (m Maybe3[T0, T1, T2]) Error() error {
	return m.Err
}

// HasError returns true if there's an error, false otherwise
func (m Maybe3[T0, T1, T2]) HasError() bool {
	return m.Err != nil
}
