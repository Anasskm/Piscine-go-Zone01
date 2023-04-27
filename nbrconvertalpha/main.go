package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	a := os.Args[1:]
	z := "abcdefghijklmnopqrstuvwxyz"
	flag := false
	for _, i := range a {
		if i == "--upper" {
			a = a[1:]
			flag = true
		}
	}

	for _, i := range a {
		index := Atoi(i)
		if index <= 26 && index >= 1 {
			if flag {
				z01.PrintRune(rune(z[index-1] - 32))
			} else {
				z01.PrintRune(rune(z[index-1]))
			}
		} else {
			z01.PrintRune(' ')
		}

	}

	z01.PrintRune('\n')
}

func Atoi(s string) int {
	x := 0

	for _, i := range s {

		if i < '0' || i > '9' {
			return 0
		}
		y := 0
		for k := '1'; k <= i; k++ {
			y++
		}
		x = x*10 + y
	}
	return x
}
