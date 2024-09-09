package results

// True is the interface for [True1] and [True2].
type True interface {
	// Values returns the stored values.
	Values() []any

	// IsOK returns the value of OK.
	IsOK() bool
}
