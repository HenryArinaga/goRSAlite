package main

import (
	"bufio"
	"fmt"
	"goRSAlite/factorization"
	"goRSAlite/rsa"
	"os"
	"time"
)

var number_to_factor int
var factor_method int

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var factor_Result []int

	fmt.Println("RSA Key Generation and Factorization Tool")
	fmt.Println("This program will factor a number and compute RSA keys based on the factors.")

	fmt.Println("Enter a message to encrypt:")

	var input string
	scanner.Scan()
	input = scanner.Text()
	message := []byte(input)

	fmt.Println("Enter a number to use as a RSA factor:")
	fmt.Scanln(&number_to_factor)

	for i := 0; i < len(message); i++ {
		if int(message[i]) >= number_to_factor || int(message[i]) <= 0 {
			fmt.Printf("Message byte %d must satisfy 0 < byte < n for RSA encryption.\n", message[i])
			return
		}
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
	case 3:
		factor_Result = factorization.FactorFermant(number_to_factor)
	default:
		fmt.Println("Invalid method selected")
	}
	start := time.Now()
	rsa_message(factor_Result, int(message[0]), input)
	duration := time.Since(start)
	fmt.Printf("\nExecution time: %s\n", duration)
}

func rsa_message(factor_Result []int, message int, input string) {

	var decrypted_message []int
	var ciphertext []int
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

			for i := 0; i < len(input); i++ {
				ciphertext = append(ciphertext, rsa.Encrypt(int(input[i]), e, p*q))
			}
			fmt.Printf("Number of encrypted bytes: %d\n", len(ciphertext))
			fmt.Printf("Ciphertext: %v \n", ciphertext)

			for i := 0; i < len(ciphertext); i++ {
				decrypted_message = append(decrypted_message, rsa.Decrypt(ciphertext[i], d, p*q))
			}
			fmt.Printf("Number of decrypted bytes: %d\n", len(decrypted_message))
			fmt.Printf("Decrypted Message: ")
			for i := 0; i < len(decrypted_message); i++ {
				fmt.Printf("%c", decrypted_message[i])
			}

		}

	}
}
