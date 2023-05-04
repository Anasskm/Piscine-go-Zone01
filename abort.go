package piscine

func Abort(a, b, c, d, e int) int {
	var s []int

	s = append(s, a, b, c, d, e)
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] < s[j] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
	return s[2]
}
