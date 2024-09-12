package results

// TypeAs returns the result of type assertions.
// This is a function, so it is addressable.
//
//	value, ok := v.(T)
func TypeAs[T any](v any) (value T, ok bool) {
	value, ok = v.(T)
	return
}
