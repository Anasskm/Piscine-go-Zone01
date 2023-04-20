package piscine

func SortIntegerTable(t []int) {
	for i := 0; i < len(t); i++ {
		for j := 0; j < len(t); j++ {
			if t[i] < t[j] {
				a := t[j]
				t[j] = t[i]
				t[i] = a
			}
		}
	}
}
