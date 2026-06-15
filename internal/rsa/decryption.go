package rsa

func Decrypt(ciphertext int, d int, number_to_factor int) int {

	base := ciphertext
	result := 1

	for d > 0 {
		if d%2 != 0 {
			result = (result * base) % number_to_factor
		}
		d = d / 2
		base = (base * base) % number_to_factor
	}
	return result
}
