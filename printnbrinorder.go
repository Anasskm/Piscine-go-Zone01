package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	}
	for _, j := range SortTable(Convtotable(n)) {
		j += '0'

		z01.PrintRune(rune(j))
	}
}

func Convtotable(n int) (s []int) {
	for n > 0 {

		s = append(s, n%10)

		n /= 10

	}
	return s
}

func SortTable(t []int) []int {
	for i := 0; i < len(t); i++ {
		for j := 0; j < len(t); j++ {
			if t[i] < t[j] {
				a := t[j]
				t[j] = t[i]
				t[i] = a
			}
		}
	}
	return t
}
