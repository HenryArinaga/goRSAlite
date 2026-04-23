package rsa

// GCD computes the greatest common divisor of two integers using the Euclidean algorithm.
// Here we are using recursion
func GCD(a int, b int) int {
	if b == 0 {
		return a
	}
	// The GCD of a and b is the same as the GCD of b and the remainder of a divided by b.
	return GCD(b, a%b)
}
