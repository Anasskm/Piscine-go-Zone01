package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)
	x := "x = "
	y := " y = "

	rx := []rune{}
	ry := []rune{}
	for points.x != 0 {
		rx = append(rx, rune(points.x%10)+'0')
		points.x /= 10
	}
	for points.y != 0 {
		ry = append(ry, rune(points.y%10)+'0')
		points.y /= 10
	}

	for _, i := range x {
		z01.PrintRune(rune(i))
	}
	for i := len(rx) - 1; i >= 0; i-- {
		z01.PrintRune(rx[i])
	}
	z01.PrintRune(',')
	for _, i := range y {
		z01.PrintRune(rune(i))
	}
	for i := len(ry) - 1; i >= 0; i-- {
		z01.PrintRune(ry[i])
	}
}
