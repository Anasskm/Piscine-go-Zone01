package piscine

func BasicAtoi2(s string) int {
	x := 0
	for _, i := range s {
		if i < '0' || i > '9' {
			return 0
		}
		y := 0
		for j := '1'; j <= i; j++ {
			y++
		}
		x = x*10 + y
	}
	return x
}
