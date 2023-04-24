package piscine

func IsPrime(nb int) bool {
	x := true
	for i := 2; i < nb; i++ {
		if nb%i == 0 && nb != 2 {
			x = false
			break

		}
	}
	return x
}
