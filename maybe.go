package results

type Maybe interface {
	// Values returns the stored values.
	Values() []any

	// Error returns the error
	Error() error

	// HasError returns true if there's an error, false otherwise
	HasError() bool
}
