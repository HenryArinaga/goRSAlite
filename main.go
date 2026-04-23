package main

import (
	"fmt"
	"goRSAlite/factorization"
	"goRSAlite/rsa"
)

var n int
var factor_method int

func main() {

	fmt.Println("Enter a number to factor:")
	fmt.Scanln(&n)

	fmt.Println("Enter 1 for Trial Division, 2 for Sqrt Method:")
	fmt.Scanln(&factor_method)

	var factor_Result []int

	switch factor_method {
	case 1:
		factor_Result = factorization.FactorTrialDivison(n)
	case 2:
		factor_Result = factorization.FactorSqrt(n)
	default:
		fmt.Println("Invalid method selected")
	}
	fmt.Printf("Factors of %d: %v \n", n, factor_Result)

	if len(factor_Result) < 2 {
		fmt.Println("Not enough factors to compute RSA keys.")
		return
	} else if len(factor_Result) > 2 {
		fmt.Println("Number is not a product of two primes, cannot compute RSA keys.")
		return
	} else {
		p := factor_Result[0]
		q := factor_Result[1]
		if p == q {
			fmt.Println("Both factors are the same, number is not a product of" +
				"" + "two distinct primes, cannot compute RSA keys.")
			return
		} else {
			fmt.Printf("p: %d, q: %d \n", p, q)
			fmt.Printf("Euler's Totient: %d \n", rsa.Totient(p, q))
		}

	}
}
