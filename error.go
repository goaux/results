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
