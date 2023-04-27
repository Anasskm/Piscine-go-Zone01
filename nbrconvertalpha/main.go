package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	a := os.Args[1:]
	z := "abcdefghijklmnopqrstuvwxyz"

	if a[0] == "--upper" {
		for i := 1; i < len(a); i++ {
			index := Atoi(a[i])
			if index > 26 || index <= 0 {
				z01.PrintRune(' ')
			} else {
				alpha := z[index-1]
				z01.PrintRune(rune(alpha - 32))
			}

		}
	} else {
		for i := 0; i < len(a); i++ {
			index := Atoi(a[i])
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

func Atoi(s string) int {
	x := 0
	si := 1
	for j, i := range s {
		if s[j] == '-' && j == 0 {
			si = -1
			continue
		}
		if s[j] == '+' && j == 0 {
			continue
		}
		if i < '0' || i > '9' {
			return 0
		}
		y := 0
		for k := '1'; k <= i; k++ {
			y++
		}
		x = x*10 + y
	}
	return x * si
}
