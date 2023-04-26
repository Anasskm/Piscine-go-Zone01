package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	r := os.Args[0]
	var k int

	for i := 0; i < len(r); i++ {
		if r[len(r)-1-i] == '/' {
			k = len(r) - 1 - i
			break
		}
	}
	s := r[k+1:]
	for _, i := range s {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
