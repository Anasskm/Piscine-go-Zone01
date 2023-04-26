package piscine

func BasicJoin(e []string) string {
	var st string
	for i := range e {
		r := []rune(e[i])
		s := string(r)
		st += s
	}
	return st
}
