package piscine

func Map(f func(int) bool, a []int) []bool {
	var b []bool
	for _, i := range a {
		b = append(b, f(i))
	}
	return b
}
