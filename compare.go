package piscine

func Compare(a, b string) int {
	ra := []rune(a)
	rb := []rune(b)
	if len(ra) == len(rb) {
		count := 0
		for i := range ra {
			if ra[i] == rb[i] {
				count++
			}
		}
		if count != len(ra) {
			return 1
		}
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
