package factorization

func IsPrime(number_to_factor int) bool {
	if number_to_factor < 2 {
		return false

	}

	for i := 2; i*i <= number_to_factor; i++ {
		if number_to_factor%i == 0 {
			return false
		}
	}

	return true

}
