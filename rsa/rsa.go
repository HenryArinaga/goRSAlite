/*
This package implements the core functions of the RSA algorithm,
it finds the totient of p and q, gcd of E and D
*/
package rsa

func Totient(p int, q int) int {

	return (p - 1) * (q - 1)

}

/*
	GCD computes the greatest common divisor

of two integers using the Euclidean algorithm.
Here we are using recursion
*/
func GCD(a int, b int) int {
	if b == 0 {
		return a
	}
	// The GCD of a and b is the same as the
	//GCD of b and the remainder of a divided by b.
	return GCD(b, a%b)
}

func E(totient int) int {
	//start at 3 because 1 and 2 are not valid choices for e
	for e := 3; e < totient; e++ {
		//take the GCD of e and the totient,
		// if it is 1 then we have found a valid e value
		if GCD(e, totient) == 1 {
			return e
		}
	}
	// retrn -1 if no valid e value is found,
	// which should not happen for reasonable totient values
	return -1
}

// D function computes the private exponent
// d given the public exponent e and the totient.
func D(e int, totient int) int {
	for d := 1; d < totient; d++ {
		if (d*e)%totient == 1 {
			return d
		}
	}
	return -1
}

/* add euclid method here for computing d,
which is more efficient than the brute-force method above */
