//go:build go1.21

package results_test

import "slices"

func slicesEqual[S ~[]E, E comparable](s1, s2 S) bool {
	return slices.Equal(s1, s2)
}
