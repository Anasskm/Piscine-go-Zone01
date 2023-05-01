package main

import "github.com/01-edu/z01"

func setPoint(p *[2]int) {
	p[0] = 42
	p[1] = 21
}

func main() {
	var point [2]int
	points := &point

	setPoint(points)

	p0 := string(point[0])
	p1 := string(point[1])
	r1 := []rune(p0)
	r2 := []rune(p1)
	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	for _, i := range r1 {
		z01.PrintRune(i)
	}
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	for _, i := range r2 {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
