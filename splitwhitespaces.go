package piscine

func SplitWhiteSpaces(s string) []string {
	k := 1

	word := ""
	for i, j := range s {
		if j == ' ' && s[i+1] != ' ' {
			k++
		}
	}
	result := make([]string, k)
	index := 0
	for _, y := range s {
		if y != ' ' {
			word += string(y)
		} else {
			result[index] = word
			word = ""
			index++
		}
	}
	result[index] = word
	return result
}
