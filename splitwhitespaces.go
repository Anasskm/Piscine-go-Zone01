package piscine

func SplitWhiteSpaces(s string) []string {
	word := ""

	var result []string

	for _, y := range s {
		if y != ' ' && y != '\t' {
			word += string(y)
		} else {
			result = append(result, word)
			word = ""

		}
	}
	result = append(result, word)
	return result
}
