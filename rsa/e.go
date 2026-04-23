package rsa

func E(totient int) int {
	//start at 3 because 1 and 2 are not valid choices for e
	for e := 3; e < totient; e++ {
		//take the GCD of e and the totient,
		// if it is 1 then we have found a valid e value
		if GCD(e, totient) == 1 {
			return e
		}
	}
	return -1
}
