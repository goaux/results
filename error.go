package results

import "errors"

// ErrorAs returns the result of calling [errors.As](err, &target).
//
// The coding pattern using [errors.As] is following:
//
//	var pathError *fs.PathError
//	if errors.As(err, &pathError) {
//
// Using ErrorAs is a bit more concise:
//
//	if pathError, ok := results.ErrorAs[*fs.PathError](err); ok {
//
// Improvements:
//
//   - Just one line.
//   - Create a scope for the pathError variable.
func ErrorAs[T any](err error) (target T, ok bool) {
	ok = errors.As(err, &target)
	return
}

// ErrorIs returns true if the error or any error in its chain
// matches any of the target errors. It returns false otherwise.
//
// It iterates through the target errors and uses [errors.Is] to
// determine if the input error or any error in its chain matches
// any of them.
//
// Parameters:
//
//	err: The error to check.
//	target: A variadic list of target errors to compare against.
//
// Returns:
//
//	True if the error matches any of the target errors, false otherwise.
func ErrorIs(err error, target ...error) bool {
	for _, e := range target {
		if errors.Is(err, e) {
			return true
		}
	}
	return false
}
