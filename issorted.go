package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	var z bool
	if f(a[0], a[1]) < 0 {
		for i := 1; i < len(a); i++ {
			for j := i + 1; j < len(a); j++ {
				if f(a[i], a[j]) > 0 {
					z = false
					break
				}
			}
		}

		z = true
	}
	if f(a[0], a[1]) > 0 {
		for i := 1; i < len(a); i++ {
			for j := i + 1; j < len(a); j++ {
				if f(a[i], a[j]) < 0 {
					z = false
					break
				}
			}
		}

		z = true
	}

	return z
}
