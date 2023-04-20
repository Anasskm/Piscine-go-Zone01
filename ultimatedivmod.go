package piscine

func UltimateDivMod(a *int, b *int) {
	var c int
	var d int
	c = *a / *b
	d = *a % *b
	*a = c
	*b = d
}
