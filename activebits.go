package piscine

func ActiveBits(n int) uint {
	if n < 2 {
		return uint(n)
	}
	return (uint(n) % 2) + ActiveBits(n/2)
}
