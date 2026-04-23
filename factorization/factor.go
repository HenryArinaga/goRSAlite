package factorization

func FactorTrialDivison(n int) []int {

	factor_array := make([]int, 0)
	for i := 2; i <= n; i++ {
		for n%i == 0 {
			factor_array = append(factor_array, i)
			n = n / i
		}
	}
	return factor_array
}

func FactorSqrt(n int) []int {

	factor_array := make([]int, 0)
	for i := 2; i <= (n*n)/2; i++ {
		for n%i == 0 {
			factor_array = append(factor_array, i)
			n = n / i
		}
	}
	return factor_array
}

/*func FactorFermant(n int) []int {

	factor_array := make ([]int, 0)

	for i := 1; i < n; i++ {

	}


} */
