package main

import (
	"bufio"
	"context"
	"fmt"
	"goRSAlite/internal/factorization"
	"goRSAlite/internal/rsa"
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

	fmt.Println("Enter 1 for Trial Division, 2 for Sqrt Method, 3 for Fermat's Method:")
	fmt.Scanln(&factor_method)
	if factor_method != 1 && factor_method != 2 && factor_method != 3 {
		fmt.Println("Invalid method selected, please enter 1, 2, or 3.")
		return
	}
	factorResult := make(chan []int, 1)
	ctx := context.Background()

	switch factor_method {
	case 1:
		factorization.FactorTrialDivision(number_to_factor, factorResult, ctx)
		factor_Result = <-factorResult

	case 2:
		factorization.FactorSqrt(number_to_factor, factorResult, ctx)
		factor_Result = <-factorResult
	case 3:
		factor_Result = factorization.FactorFermat(number_to_factor, ctx)
	default:
		fmt.Println("Invalid method selected")
	}

	rsa_message(factor_Result, input)

}

func rsa_message(factor_Result []int, input string) {

	var decrypted_message []int
	var ciphertext []int
	start := time.Now()
	fmt.Printf("Factors of %d: %v \n", number_to_factor, factor_Result)
	duration := time.Since(start)
	fmt.Printf("\nTime to find Factors: %s\n", duration)

	if len(factor_Result) < 2 {
		fmt.Println("Not enough factors to compute RSA keys.")
		return
	} else if len(factor_Result) > 2 {
		fmt.Println("Number is not a product of two primes, cannot compute RSA keys.")
		return
	} else {
		fmt.Println("Enter 1 to compute RSA keys and encrypt the message, 2 to skip:")
		var calculate_rsa int
		fmt.Scanln(&calculate_rsa)
		switch calculate_rsa {
		case 1:
			start := time.Now()
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
			duration := time.Since(start)
			fmt.Printf("\nExecution time: %s\n", duration)
		case 2:
			fmt.Println("RSA key generation and encryption skipped.")
		default:
			fmt.Println("Invalid option selected, skipping RSA key generation and encryption.")
		}
	}
}
