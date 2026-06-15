package rsa

func Encrypt(message int, e int, number_to_factor int) int {

	base := message
	result := 1

	for e > 0 {
		if e%2 != 0 {
			result = (result * base) % number_to_factor
		}
		e = e / 2
		base = (base * base) % number_to_factor
	}
	return result
}
