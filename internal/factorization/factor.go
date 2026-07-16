package factorization

import (
	"context"
	"fmt"
	"math"
)

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

func FactorFermat(number_to_factor int, ctx context.Context) []int {

	factor_array := make([]int, 0)
	select {
	case <-ctx.Done():
		fmt.Println("Ctx cancelled")
		return []int{}
	default:
	}

	a := int(math.Ceil(math.Sqrt(float64(number_to_factor))))
	var bSquared int

	if number_to_factor <= 1 {
		return []int{}
	}

	if IsPrime(number_to_factor) {
		return []int{number_to_factor}
	}

	if number_to_factor%2 == 0 {
		return append([]int{2}, FactorFermat(number_to_factor/2, ctx)...)
	}

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Ctx cancelled")
			return []int{}
		default:
		}

		bSquared = a*a - number_to_factor

		sqrt_b := math.Sqrt(float64(bSquared))
		if sqrt_b == float64(int(sqrt_b)) {
			left := a - int(sqrt_b)
			right := a + int(sqrt_b)

			leftFactors := FactorFermat(left, ctx)
			rightFactors := FactorFermat(right, ctx)
			factor_array = append(factor_array, leftFactors...)
			factor_array = append(factor_array, rightFactors...)
			break
		}
		a = a + 1

	}

	return factor_array
}
