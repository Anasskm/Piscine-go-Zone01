package piscine

func Compare(a, b string) int {
	ra := []rune(a)
	rb := []rune(b)
	if len(a) == len(b) {
		return 0
	}
	if ra[0] != rb[0] {
		return -1
	}
	if ra[len(ra)-1] != rb[len(rb)-1] {
		return 1
	}
	return 0
}
