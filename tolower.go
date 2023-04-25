package piscine

func ToLower(s string) string {
	r := []rune(s)

	for i := range r {
		if r[i] >= 65 && r[i] <= 90 {
			r[i] += 32
		}
	}
	st := string(r)
	return st
}
