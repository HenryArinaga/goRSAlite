package main

import (
	"fmt"
	"goRSAlite/factorization"
	"goRSAlite/rsa"
)

var number_to_factor int
var factor_method int

func main() {

	var factor_Result []int

	fmt.Println("RSA Key Generation and Factorization Tool")
	fmt.Println("This program will factor a number and compute RSA keys based on the factors.")

	fmt.Println("Enter a integer message to encrypt:")

	var message int
	fmt.Scanln(&message)

	fmt.Println("Enter a number to factor:")
	fmt.Scanln(&number_to_factor)
	if message >= number_to_factor || message <= 0 {
		fmt.Println("Message must satisfy 0 < message < n for RSA encryption.")
		return
	}

	fmt.Println("Enter 1 for Trial Division, 2 for Sqrt Method:")
	fmt.Scanln(&factor_method)
	if factor_method != 1 && factor_method != 2 {
		fmt.Println("Invalid method selected, please enter 1 or 2.")
		return
	}

	switch factor_method {
	case 1:
		factor_Result = factorization.FactorTrialDivison(number_to_factor)
	case 2:
		factor_Result = factorization.FactorSqrt(number_to_factor)
	default:
		fmt.Println("Invalid method selected")
	}
	rsa_message(factor_Result, message)

}

func rsa_message(factor_Result []int, message int) {

	fmt.Printf("Factors of %d: %v \n", number_to_factor, factor_Result)

	if len(factor_Result) < 2 {
		fmt.Println("Not enough factors to compute RSA keys.")

		return
	} else if len(factor_Result) > 2 {
		fmt.Println("Number is not a product of two primes, cannot compute RSA keys.")
		return
	} else {
		p := factor_Result[0]
		q := factor_Result[1]
		totient := rsa.Totient(p, q)
		e := rsa.E(totient)
		d := rsa.D(e, totient)
		if p == q {
			fmt.Println("Both factors are the same, number is not a product of" +
				"" + "two distinct primes, cannot compute RSA keys.")
			return
		} else {
			fmt.Printf("p: %d, q: %d \n", p, q)
			fmt.Printf("Euler's Totient: %d \n", totient)
			fmt.Printf("Public Exponent e: %d \n", e)
			fmt.Printf("Private Exponent d: %d \n", d)
			fmt.Printf("Public Key: (%d, %d)\n", e, p*q)
			fmt.Printf("Private Key: (%d, %d)\n", d, p*q)
			ciphertext := rsa.Encrypt(message, e, p*q)
			fmt.Printf("Ciphertext: %d \n", ciphertext)
			decrypted_message := rsa.Decrypt(ciphertext, d, p*q)
			fmt.Printf("Decrypted Message: %d \n", decrypted_message)
		}

	}
}
