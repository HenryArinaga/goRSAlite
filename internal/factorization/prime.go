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

func Sieve(number_to_factor int) []int {
	var res []int
	isPrime := make([]bool, number_to_factor+1)
	for i := 2; i <= number_to_factor; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i <= number_to_factor; i++ {
		if isPrime[i] == true {
			for multiple := i * i; multiple <= number_to_factor; multiple += i {
				isPrime[multiple] = false
			}
		}
	}
	for i := 2; i <= number_to_factor; i++ {
		if isPrime[i] == true {
			res = append(res, i)
		}
	}
	return res
}
