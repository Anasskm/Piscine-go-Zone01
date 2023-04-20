package piscine

import "fmt"

func PrintStr(s string) {
	x := []byte(s)
	for i := 0; i < len(s); i++ {
		fmt.Print(string(x[i]))
	}
}
