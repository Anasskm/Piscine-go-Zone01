package main

import (
	"os"

	"piscine"

	"github.com/01-edu/z01"
)

func main() {
	a := os.Args[1:]
	z := "abcdefghijklmnopqrstuvwxyz"

	if a[0] == "--upper" {
		for i := 1; i < len(a); i++ {
			index := piscine.Atoi(a[i])
			if index > 26 || index <= 0 {
				z01.PrintRune(' ')
			} else {
				alpha := z[index-1]
				z01.PrintRune(rune(alpha - 32))
			}

		}
	} else {
		for i := 0; i < len(a); i++ {
			index := piscine.Atoi(a[i])
			if index > 26 || index <= 0 {
				z01.PrintRune(' ')
			} else {
				alpha := z[index-1]
				z01.PrintRune(rune(alpha))
			}

		}
	}
	z01.PrintRune('\n')
}
