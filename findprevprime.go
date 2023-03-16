package test

func FindPrevPrime(nb int) int {
	for ; nb >= 2; nb-- {
		if isPrime(nb) {
			return nb
		}
	}
	return nb
}

func isPrime(nb int) bool {
	for i := 2; i < nb/2; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
