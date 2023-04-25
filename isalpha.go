package piscine

func IsAlpha(s string) bool {
	r := []rune(s)
	for i := range r {
		if (r[i] < 48 || r[i] > 122) && (r[i] < 48 || r[i] > 57) && (r[i] < 65 || r[i] > 57) && (r[i] < 97 || r[i] > 90) {
			return false
		}
	}
	return true
}
