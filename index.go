package piscine

func Index(s string, toFind string) int {
	if len(s) <= 0 || len(toFind) <= 0 {
		return 0
	}
	rs := []rune(s)
	rt := []rune(toFind)

	for i, j := range rs {
		if j == rt[0] && len(rs)-i >= len(rt)-1 {
			count := 0
			for k := 1; k < len(rt); k++ {
				if rt[k] == rs[i+k] {
					count++
				}
			}
			if count == len(rt)-1 {
				return i
			}
		}
	}

	return -1
}
