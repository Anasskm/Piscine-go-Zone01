package piscine

func Index(s string, toFind string) int {
	rs := []rune(s)
	rt := []rune(toFind)
	for i := range rs {
		if rs[i] == rt[0] {
			return i
		}
	}
	return -1
}
