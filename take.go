package results

// Take2 returns the first of the two provided arguments, discarding the second.
func Take2[T0, T1 any](a T0, _ T1) T0 { return a }

// Drop2 discards the first of the two provided arguments and returns the second.
func Drop2[T0, T1 any](_ T0, b T1) T1 { return b }

// Take3 returns the first of the three provided arguments, discarding the other two.
func Take3[T0, T1, T2 any](a T0, _ T1, _ T2) T0 { return a }

// Drop3 discards the first of the three provided arguments and returns the remaining two.
func Drop3[T0, T1, T2 any](_ T0, b T1, c T2) (T1, T2) { return b, c }
