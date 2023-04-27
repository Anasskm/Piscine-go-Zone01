package piscine

func ConcatParams(args []string) string {
	var s string
	for i, j := range args {
		s += j
		if i != len(args)-1 {
			s += "\n"
		}
	}
	return s
}
