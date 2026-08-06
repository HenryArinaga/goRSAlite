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
