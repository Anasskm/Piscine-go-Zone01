package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	a := os.Args
	for i := 1; i < len(a); i++ {
		for j := range a[i] {
			r := []rune(a[i])
			if r[j] >= '0' && r[j] <= '9' {
				z01.PrintRune(r[j])
				z01.PrintRune('\n')
			}
		}
	}
	for i := 1; i < len(a); i++ {
		for j := range a[i] {
			r := []rune(a[i])
			if r[j] >= 'A' && r[j] <= 'Z' {
				z01.PrintRune(r[j])
				z01.PrintRune('\n')
			}
		}
	}
	for i := 1; i < len(a); i++ {
		for j := range a[i] {
			r := []rune(a[i])
			if r[j] >= 'a' && r[j] <= 'z' {
				z01.PrintRune(r[j])
				z01.PrintRune('\n')
			}
		}
	}
}
