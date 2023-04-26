package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	a := os.Args

	for i := range a {
		if i != 0 {
			for j := range a[i] {
				runes := []rune(a[i])
				{
					z01.PrintRune(runes[j])
				}
			}
			z01.PrintRune('\n')
		}
	}
}
