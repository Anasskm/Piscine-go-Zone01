package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}
	x := 1
	for i := 1; i <= power; i++ {
		x *= nb
	}
	return x
}
