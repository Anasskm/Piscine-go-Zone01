package piscine

func IsPrime(nb int) bool {
	x := true
	if nb < 2 {
		return false
	}
	for i := 2; i < nb; i++ {
		if nb%i == 0 && nb != 2 {
			x = false
			break

		}
		x = true
	}
	return x
}
