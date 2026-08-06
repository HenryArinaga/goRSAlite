package rsa

// m = c^d % n

func mulMod(a, b, mod int64) int64 {
	var result int64
	a = a % mod

	for b > 0 {
		if b%2 == 1 {
			result = (result + a) % mod
		}
		a = (a + a) % mod
		b = b / 2
	}

	return result
}

func powMod(base, exp, mod int64) int64 {
	var result int64 = 1
	base = base % mod

	for exp > 0 {
		if exp%2 == 1 {
			result = mulMod(result, base, mod)
		}
		base = mulMod(base, base, mod)
		exp = exp / 2
	}

	return result
}

func Decrypt(ciphertext int, d int, number_to_factor int) int {
	return int(powMod(int64(ciphertext), int64(d), int64(number_to_factor)))
}
