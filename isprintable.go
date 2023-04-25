package piscine

func IsPrintable(s string) bool {
	r := []rune(s)
	for i := range r {
		if r[i] < 33 || r[i] > 127 {
			return false
		}
	}
	return true
}
