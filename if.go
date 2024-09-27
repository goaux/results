package results

// If returns the true value if ok is true, otherwise the false value.
func If[T any](ok bool, true, false T) T {
	if ok {
		return true
	} else {
		return false
	}
}

// IfFunc returns the result of calling true function if ok is true, otherwise the result of calling false function.
func IfFunc[T any](ok bool, true, false func() T) T {
	if ok {
		return true()
	} else {
		return false()
	}
}
