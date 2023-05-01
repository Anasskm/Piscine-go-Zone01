package main

import "fmt"

func setPoint(p *[2]int) {
	p[0] = 42
	p[1] = 21
}

func main() {
	var point [2]int
	points := &point

	setPoint(points)

	fmt.Printf("x = %d, y = %d\n", point[0], point[1])
}
