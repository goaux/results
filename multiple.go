package results

// Get1of2 returns the first parameter of type A from two parameters.
func Get1of2[A, B any](a A, _ B) A { return a }

// Get2of2 returns the second parameter of type B from two parameters.
func Get2of2[A, B any](_ A, b B) B { return b }

// Get1of3 returns the first parameter of type A from three parameters.
func Get1of3[A, B, C any](a A, _ B, _ C) A { return a }

// Get2of3 returns the second parameter of type B from three parameters.
func Get2of3[A, B, C any](_ A, b B, _ C) B { return b }

// Get3of3 returns the third parameter of type C from three parameters.
func Get3of3[A, B, C any](_ A, _ B, c C) C { return c }

// Get12of3 returns the first two parameters of types A and B from three parameters.
func Get12of3[A, B, C any](a A, b B, _ C) (A, B) { return a, b }

// Get23of3 returns the last two parameters of types B and C from three parameters.
func Get23of3[A, B, C any](_ A, b B, c C) (B, C) { return b, c }
