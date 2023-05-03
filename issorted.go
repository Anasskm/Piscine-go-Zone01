package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	var z bool
	for i := range a {
		for j := i + 1; j < len(a); j++ {
			if f(a[i], a[j]) > 0 {
				for x := j; x < len(a); x++ {
					for y := j + 1; y < len(a); y++ {
						if f(a[x], a[y]) < 0 {
							z = false
							break
						}
					}
				}
				z = true
				break

			}
			if f(a[i], a[j]) < 0 {
				for x := j; x < len(a); x++ {
					for y := j + 1; y < len(a); y++ {
						if f(a[x], a[y]) > 0 {
							z = false
							break
						}
					}
				}
				z = true
				break

			}
		}
	}
	return z
}
