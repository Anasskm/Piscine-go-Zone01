package piscine

func Sqrt(nb int) int {
	if nb < 0 {
		return 0
	}

	x := 0

	for i := 0; i < nb; i++ {
		if i*i == nb {
			x = i
			break
		}
	}

	return x
}
