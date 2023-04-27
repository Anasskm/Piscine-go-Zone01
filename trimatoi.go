package piscine

func TrimAtoi(s string) int {
	x := 0
	si := 1

	if Numeric(s) {
		for _, i := range s {
			if i < '0' || i > '9' {
				continue
			}
			y := 0
			for k := '1'; k <= i; k++ {
				y++
			}
			x = x*10 + y
		}
		for j := range s {

			if s[j] >= '0' && s[j] <= '9' {
				break
			}
			if s[j] == '-' {
				si = -1
			}

		}
		return x * si
	} else {
		return 0
	}
}

func Numeric(s string) bool {
	r := []rune(s)
	tr := false
	for i := range r {
		if r[i] >= 48 && r[i] <= 57 {
			tr = true
			break
		}
	}
	return tr
}
