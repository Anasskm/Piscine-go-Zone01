package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	a := os.Args[1:]
	flag := false
	for i := range a {
		if i == 0 && a[i] == "--upper" {
			a = a[1:]
			flag = true
			break
		}
	}
	for _, i := range a {
		digit := 0
		for _, b := range i {
			digit = digit*10 + int(b-'0')
		}
		if digit <= 26 && digit >= 1 {
			if flag {
				z01.PrintRune('a' + rune(digit-32) - 1)
			} else {
				z01.PrintRune('a' + rune(digit) - 1)
			}
		} else {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}
