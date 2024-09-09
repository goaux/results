package results

// True2 holds values and a bool.
type True2[T0, T1 any] struct {
	V0 T0
	V1 T1
	OK bool
}

var _ True = True2[any, any]{}

// If2 creates a [True2] from v0, v1 and ok.
func If2[T0, T1 any](v0 T0, v1 T1, ok bool) True2[T0, T1] {
	return True2[T0, T1]{V0: v0, V1: v1, OK: ok}
}

// Or returns t.V0 and t.V1 if t.OK = true, otherwise v0, v1.
func (t True2[T0, T1]) Or(v0 T0, v1 T1) (T0, T1) {
	if t.OK {
		return t.V0, t.V1
	}
	return v0, v1
}

// OrFunc returns t.V0 and t.V1 if t.OK = true, otherwise the result of calling f.
func (t True2[T0, T1]) OrFunc(f func() (T0, T1)) (T0, T1) {
	if t.OK {
		return t.V0, t.V1
	}
	return f()
}

// OrPanic returns t.V0 and t.V1 if t.OK = true, otherwise calling [panic](msg).
func (t True2[T0, T1]) OrPanic(msg any) (T0, T1) {
	if t.OK {
		return t.V0, t.V1
	}
	panic(msg)
}

// Values returns the stored value.
func (t True2[T0, T1]) Values() []any {
	return []any{t.V0, t.V1}
}

// IsOK returns t.OK.
func (t True2[T0, T1]) IsOK() bool {
	return t.OK
}
