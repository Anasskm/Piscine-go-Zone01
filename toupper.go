package piscine

func ToUpper(s string) string {
	r := []rune(s)

	for i := range r {
		if r[i] >= 97 && r[i] <= 122 {
			r[i] -= 32
		}
	}
	st := string(r)
	return st
}
