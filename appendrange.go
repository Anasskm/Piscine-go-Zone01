package piscine

func AppendRange(min, max int) []int {
	s := []int(nil)
	l := max - min
	if l <= 0 {
		return s
	}

	for i := 0; i < l; i++ {
		s = append(s, min+i)
	}
	return s
}
