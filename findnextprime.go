package piscine

func FindNextPrime(nb int) int {
	if nb < 1000079800 {
		if IsPrime(nb) {
			return nb
		} else {
			return FindNextPrime(nb + 1)
		}
	}
	return 0
}
