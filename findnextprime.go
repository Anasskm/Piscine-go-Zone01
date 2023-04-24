package piscine

func FindNextPrime(nb int) int {
	if nb <= 2 {
		return 2
	}
	if nb%2 == 0 {
		nb = nb + 1
	}
	for {
		if Ispe(nb) {
			return nb
		}
		nb = nb + 2
	}
}

func Ispe(nb int) bool {
	if nb < 2 {
		return false
	}
	if nb%2 == 0 {
		return false
	}
	for i := 3; i <= nb/3; i = i + 2 {
		if nb%i == 0 {
			return false
		}
	}

	return true
}
