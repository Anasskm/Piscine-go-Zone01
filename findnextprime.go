package piscine

func FindNextPrime(nb int) int {
	if nb < 0 {
		return 2
	}
	if IsPrime(nb) {
		return nb
	} else {
		for i := nb; i <= nb+100; i++ {
			if IsPrime(i) {
				return i
			}
		}
	}
	return 0
}
