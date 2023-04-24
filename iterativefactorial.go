package piscine

func IterativeFactorial(nb int) int {
	x := 1
	if nb < 0 || nb > 25 {
		return 0
	} else {
		for i := 1; i < nb+1; i++ {
			x *= i
		}
	}
	return x
}
