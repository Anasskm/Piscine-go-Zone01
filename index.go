package piscine

func Index(s string, toFind string) int {
	if len(s) > 0 && len(toFind) > 0 {
		rs := []rune(s)

		rt := []rune(toFind)
		for i := range rs {
			if rs[i] == rt[0] {
				return i
			}
		}
		return -1
	} else {
		return -1
	}
}
