package factorization

import (
	"context"
	"fmt"
	"math"
)

func FactorTrialDivision(number_to_factor int, factorResult chan []int, ctx context.Context) {

	factor_array := make([]int, 0)
	for i := 2; i <= number_to_factor; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("Ctx cancelled")
			return
		default:
		}
		for number_to_factor%i == 0 {
			select {
			case <-ctx.Done():
				fmt.Println("Ctx cancelled")
				return
			default:
			}
			factor_array = append(factor_array, i)
			number_to_factor = number_to_factor / i
		}
	}
	factorResult <- factor_array
}

func FactorTrialDivisionSieve(number_to_factor int, factorResult chan []int, ctx context.Context) {
	sqrtLimit := int(math.Sqrt(float64(number_to_factor)))
	factor_array := make([]int, 0)
	factorSlice := Sieve(sqrtLimit)
	for _, numberToCheck := range factorSlice {
		select {
		case <-ctx.Done():
			fmt.Println("Ctx cancelled")
			return
		default:
		}
		for number_to_factor%numberToCheck == 0 {
			select {
			case <-ctx.Done():
				fmt.Println("Ctx cancelled")
				return
			default:
			}
			factor_array = append(factor_array, numberToCheck)
			number_to_factor = number_to_factor / numberToCheck
		}

	}
	if number_to_factor > 1 {
		factor_array = append(factor_array, number_to_factor)
	}
	factorResult <- factor_array
}

func FactorTrialDivisionSIMD(numberToFactor int, factorResult chan []int, ctx context.Context) {
	factorArray := make([]int, 0)

	for numberToFactor%2 == 0 {
		select {
		case <-ctx.Done():
			fmt.Println("Ctx cancelled")
			return
		default:
		}

		factorArray = append(factorArray, 2)
		numberToFactor /= 2
	}

	for i := 3; i*i <= numberToFactor; i += 8 {
		select {
		case <-ctx.Done():
			fmt.Println("Ctx cancelled")
			return
		default:
		}

		batch := [4]int{i, i + 2, i + 4, i + 6}

		for _, divisor := range batch {
			if divisor*divisor > numberToFactor {
				break
			}

			for numberToFactor%divisor == 0 {
				select {
				case <-ctx.Done():
					fmt.Println("Ctx cancelled")
					return
				default:
				}

				factorArray = append(factorArray, divisor)
				numberToFactor /= divisor
			}
		}
	}

	if numberToFactor > 1 {
		factorArray = append(factorArray, numberToFactor)
	}

	factorResult <- factorArray
}

func FactorTrialDivisionSieveSIMD(numberToFactor int, factorResult chan []int, ctx context.Context) {
	sqrtLimit := int(math.Sqrt(float64(numberToFactor)))
	factorArray := make([]int, 0)
	factorSlice := Sieve(sqrtLimit)

	for i := 0; i < len(factorSlice); i += 4 {
		select {
		case <-ctx.Done():
			fmt.Println("Ctx cancelled")
			return
		default:
		}

		end := i + 4
		if end > len(factorSlice) {
			end = len(factorSlice)
		}

		batch := factorSlice[i:end]

		for _, divisor := range batch {
			if divisor*divisor > numberToFactor {
				break
			}

			for numberToFactor%divisor == 0 {
				select {
				case <-ctx.Done():
					fmt.Println("Ctx cancelled")
					return
				default:
				}

				factorArray = append(factorArray, divisor)
				numberToFactor /= divisor
			}
		}

		if numberToFactor == 1 {
			break
		}
	}

	if numberToFactor > 1 {
		factorArray = append(factorArray, numberToFactor)
	}

	factorResult <- factorArray
}

func FactorSqrt(number_to_factor int, factorResult chan []int, ctx context.Context) {

	factor_array := make([]int, 0)
	for i := 2; i*i <= number_to_factor; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("Ctx cancelled")
			return
		default:
		}
		for number_to_factor%i == 0 {
			select {
			case <-ctx.Done():
				fmt.Println("Ctx cancelled")
				return
			default:
			}
			factor_array = append(factor_array, i)
			number_to_factor = number_to_factor / i
		}
	}
	if number_to_factor > 1 {
		factor_array = append(factor_array, number_to_factor)
	}

	factorResult <- factor_array

}

func FactorFermat(numberToFactor int, ctx context.Context) []int {
	if numberToFactor <= 1 {
		return []int{}
	}

	select {
	case <-ctx.Done():
		fmt.Println("Ctx cancelled")
		return []int{}
	default:
	}

	if IsPrime(numberToFactor) {
		return []int{numberToFactor}
	}

	// Fermat only works directly with odd numbers.
	if numberToFactor%2 == 0 {
		return append(
			[]int{2},
			FactorFermat(numberToFactor/2, ctx)...,
		)
	}

	// Remove small factors first.
	// This prevents Fermat from taking billions of iterations
	// when the factors are extremely far apart.
	smallPrimes := []int{
		3, 5, 7, 11, 13, 17, 19,
		23, 29, 31, 37, 41, 43, 47,
	}

	for _, prime := range smallPrimes {
		if numberToFactor%prime == 0 {
			return append(
				[]int{prime},
				FactorFermat(numberToFactor/prime, ctx)...,
			)
		}
	}

	a := int(math.Ceil(math.Sqrt(float64(numberToFactor))))

	// Largest int value on the current architecture.
	maxInt := int(^uint(0) >> 1)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Ctx cancelled")
			return []int{}
		default:
		}

		// Prevent a*a from overflowing.
		if a > maxInt/a {
			fmt.Println("Fermat stopped: integer overflow")
			return []int{}
		}

		bSquared := a*a - numberToFactor

		sqrtB := int(math.Sqrt(float64(bSquared)))

		// Check for a perfect square using integer arithmetic.
		if sqrtB*sqrtB == bSquared {
			left := a - sqrtB
			right := a + sqrtB

			leftFactors := FactorFermat(left, ctx)
			rightFactors := FactorFermat(right, ctx)

			factors := make([]int, 0)
			factors = append(factors, leftFactors...)
			factors = append(factors, rightFactors...)

			return factors
		}

		a++
	}
}
