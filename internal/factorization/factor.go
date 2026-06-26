package factorization

import "math"

func FactorTrialDivision(number_to_factor int) []int {

	factor_array := make([]int, 0)
	for i := 2; i <= number_to_factor; i++ {
		for number_to_factor%i == 0 {
			factor_array = append(factor_array, i)
			number_to_factor = number_to_factor / i
		}
	}
	return factor_array
}

func FactorSqrt(number_to_factor int) []int {

	factor_array := make([]int, 0)
	for i := 2; i <= (number_to_factor*number_to_factor)/2; i++ {
		for number_to_factor%i == 0 {
			factor_array = append(factor_array, i)
			number_to_factor = number_to_factor / i
		}
	}
	return factor_array
}

func FactorFermant(number_to_factor int) []int {

	factor_array := make([]int, 0)

	a := int(math.Ceil(math.Sqrt(float64(number_to_factor))))
	var bSquared int

	for {

		bSquared = a*a - number_to_factor
		if bSquared < 0 {
			continue
		}
		sqrt_b := math.Sqrt(float64(bSquared))
		if sqrt_b == float64(int(sqrt_b)) {
			factor_array = append(factor_array, a-int(sqrt_b))
			factor_array = append(factor_array, a+int(sqrt_b))
			break
		}
		a = a + 1

	}

	return factor_array
}
