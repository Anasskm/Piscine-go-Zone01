package piscine

func Sqrt(nb int) int {
	if nb < 0 {
		return 0
	}

	x := 0

	for i := 0; i < nb; i++ {
		for j := 0; j < nb; j++ {
			if i*j == nb && i == j {
				x = i
				break
			}
		}
	}
	return x
}
