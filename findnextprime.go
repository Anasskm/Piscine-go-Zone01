package piscine

func FindNextPrime(nb int) int {
	if IsPrime(nb) {
		return nb
	} else {
		for i := nb; i <= nb+10; i++ {
			if IsPrime(i) {
				return i
			}
		}
	}
	return 0
}
