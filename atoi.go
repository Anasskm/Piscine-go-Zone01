package piscine

func Atoi(s string) int {
	x := 0
	si := 1
	for j, i := range s {
		if s[j] == '-' && j == 0 {
			si = -1
			continue
		}
		if s[j] == '+' && j == 0 {
			continue
		}
		if i < '0' || i > '9' {
			return 0
		}
		y := 0
		for k := '1'; k <= i; k++ {
			y++
		}
		x = x*10 + y
	}
	return x * si
}
