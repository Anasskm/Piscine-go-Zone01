package piscine

func Unmatch(a []int) int {
	var r int
	for i := 0; i < len(a); i++ {
		r = 0
		for j := 0; j < len(a); j++ {
			if a[i] == a[j] {
				r++
			}
		}
		if r%2 != 0 {
			return a[i]
		}
	}

	return -1
}
