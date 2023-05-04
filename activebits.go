package piscine

func ActiveBits(n int) int {
	result := 1
	for ; n > 1; n = n / 2 {
		if n%2 == 1 {
			result++
		}
	}

	return result
}
