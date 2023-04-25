package piscine

func IsNumeric(s string) bool {
	r := []rune(s)
	for i := range r {
		if r[i] < 48 || r[i] > 57 {
			return false
		}
	}
	return true
}
