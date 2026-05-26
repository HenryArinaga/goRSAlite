package factorization

func FactorTrialDivison(number_to_factor int) []int {

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

/*func FactorFermant(n int) []int {

	factor_array := make ([]int, 0)

	for i := 1; i < n; i++ {

	}


} */
