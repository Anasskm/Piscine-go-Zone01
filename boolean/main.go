package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	ss := []rune(s)
	for _, r := range ss {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func main() {
	a := os.Args[1:]
	if !isEven(len(a)) {
		printStr("I have an odd number of arguments")
	} else {
		printStr("I have an even number of arguments")
	}
}

func isEven(nbr int) bool {
	if nbr%2 == 0 {
		return true
	} else {
		return false
	}
}
