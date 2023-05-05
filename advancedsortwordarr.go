package piscine

func AdvancedSortWordArr(words []string, f func(a, b string) int) {
	for i := 0; i < len(words)-1; i++ {
		for j := i + 1; j < len(words); j++ {
			if f(words[i], words[j]) == 1 {
				words[i], words[j] = words[j], words[i]
			}
		}
	}
}
