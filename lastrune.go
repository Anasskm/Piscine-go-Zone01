package piscine

func LastRune(s string) rune {
	r := []rune(s)
	return r[len(r)-1]
}
