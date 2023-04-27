package piscine

import (
	"github.com/01-edu/z01"
)

func PrintWordsTables(a []string) {
	for i := range a {
		for j := range a[i] {
			runes := []rune(a[i])
			{
				z01.PrintRune(runes[j])
			}
		}
		z01.PrintRune('\n')
	}
}
