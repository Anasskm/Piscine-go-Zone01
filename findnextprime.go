package piscine

func FindNextPrime(nb int) int {
	if nb%2 == 0 {
		nb = nb + 1
	}

	if IsPrime(nb) {
		return nb
	} else {
		return FindNextPrime(nb + 2)
	}
}
