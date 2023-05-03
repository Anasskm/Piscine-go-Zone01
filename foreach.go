package piscine

import "github.com/01-edu/z01"

func ForEach(f func(int), a []int) {
	for _, i := range a {
		f(i)
	}
	z01.PrintRune('\n')
}
