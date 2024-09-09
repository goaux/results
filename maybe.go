package results

// Maybe is the interface for [Maybe1], [Maybe2] and [Maybe3].
type Maybe interface {
	// Values returns the stored values.
	Values() []any

	// Error returns the error
	Error() error

	// HasError returns true if there's an error, false otherwise
	HasError() bool
}
