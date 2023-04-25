package piscine

func AlphaCount(s string) int {
	r := []rune(s)
	count := 0
	for i := range r {
		if r[i] >= 65 && r[i] <= 90 || r[i] >= 97 && r[i] <= 122 {
			count++
		}
	}
	return count
}
