package piscine

func Join(e []string, sep string) string {
	var st string
	for i := 0; i < len(e)-1; i++ {
		r := []rune(e[i])
		s := string(r)

		st += s + sep
	}
	rf := []rune(e[len(e)-1])
	sf := string(rf)
	st += sf
	return st
}
