package piscine

func Unmatch(a []int) int {
	var r int
	b := false
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			if a[i] == a[j] {
				r = a[i]
				b = true
			}
		}
	}
	if !b {
		return -1
	}
	return r
}
