package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	for i := range a {
		for j := i + 1; j < len(a); j++ {
			if f(a[i], a[j]) < 0 {
				return false
			}
		}
	}
	return false
}
