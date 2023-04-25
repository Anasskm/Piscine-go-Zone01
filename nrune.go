package piscine

func NRune(s string, n int) rune {
	if n < 1 || n > len(s) {
		return 0
	}
	r := []rune(s)

	return r[n-1]
}
