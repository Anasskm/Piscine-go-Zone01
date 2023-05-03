package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	if f(a[0], a[1]) <= 0 {
		for i := 1; i < len(a); i++ {
			for j := i + 1; j < len(a)-1; j++ {
				if f(a[i], a[j]) > 0 {
					return false
				}
			}
		}
		return true
	}
	if f(a[0], a[1]) > 0 {
		for i := 1; i < len(a); i++ {
			for j := i + 1; j < len(a)-1; j++ {
				if f(a[i], a[j]) < 0 {
					return false
				}
			}
		}

		return true
	}
	return false
}
