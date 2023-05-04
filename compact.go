package piscine

func Compact(ptr *[]string) int {
	var s []string
	count := 0
	for _, j := range *ptr {
		if j != "" {
			count++
			s = append(s, j)
		}
	}
	*ptr = s
	return count
}
