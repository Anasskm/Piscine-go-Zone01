package piscine

func MakeRange(min, max int) []int {
	s := []int(nil)
	l := max - min
	if l <= 0 {
		return s
	}
	s = make([]int, l)

	for i := 0; i < l; i++ {
		s[i] = min + i
	}
	return s
}
