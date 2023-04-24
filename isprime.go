package piscine

func IsPrime(nb int) bool {
	var x bool
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			x = false
			break

		}
		x = true
	}
	return x
}
