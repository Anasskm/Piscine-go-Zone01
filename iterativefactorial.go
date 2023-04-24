package piscine

func IterativeFactorial(nb int) int {
	x := 0
	if nb <= 0 {
		return x
	}
	x = 1
	for i := 1; i < nb+1; i++ {
		x = x * i
	}
	return x
}
