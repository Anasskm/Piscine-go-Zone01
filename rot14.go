package piscine

func Rot14(s string) string {
	r := []rune(s)
	for i := range r {
		if r[i] >= 'a' && r[i] <= 'z' {
			if r[i] <= 'l' {
				r[i] += 14
			} else {
				r[i] -= 12
			}
		}
		if r[i] >= 'A' && r[i] <= 'Z' {
			if r[i] <= 'L' {
				r[i] += 14
			} else {
				r[i] -= 12
			}
		}
	}
	return string(r)
}
