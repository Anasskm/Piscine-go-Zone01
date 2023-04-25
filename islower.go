package piscine

func IsLower(s string) bool {
	r := []rune(s)
	for i := range r {
		if r[i] < 97 || r[i] > 122 {
			return false
		}
	}
	return true
}
