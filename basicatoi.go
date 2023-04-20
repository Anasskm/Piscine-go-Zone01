package piscine

func BasicAtoi(s string) int {
	x := 0
	for _, i := range s {
		y := 0
		for j := '1'; j <= i; j++ {
			y++
		}
		x = x*10 + y
	}
	return x
}
