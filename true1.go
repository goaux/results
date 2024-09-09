package results

// True1 holds a value and a bool.
type True1[T any] struct {
	V0 T
	OK bool
}

var _ True = True1[any]{}

// If1 creates a [True1] from v0 and ok.
func If1[T any](v0 T, ok bool) True1[T] {
	return True1[T]{V0: v0, OK: ok}
}

// Or returns t.V0 if t.OK = true, otherwise defaultValue.
func (t True1[T]) Or(defaultValue T) T {
	if t.OK {
		return t.V0
	}
	return defaultValue
}

// OrFunc returns t.V0 if t.OK = true, otherwise the result of calling f.
func (t True1[T]) OrFunc(f func() T) T {
	if t.OK {
		return t.V0
	}
	return f()
}

// OrPanic returns t.V0 if t.OK = true, otherwise calling [panic](msg).
func (t True1[T]) OrPanic(msg any) T {
	if t.OK {
		return t.V0
	}
	panic(msg)
}

// Values returns the stored value.
func (t True1[T]) Values() []any {
	return []any{t.V0}
}

// IsOK returns t.OK.
func (t True1[T]) IsOK() bool {
	return t.OK
}
