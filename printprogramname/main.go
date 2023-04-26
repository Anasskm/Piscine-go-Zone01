package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	r := []rune(os.Args[0])
	for i := range r {
		z01.PrintRune(r[i])
	}
	z01.PrintRune('\n')
}
