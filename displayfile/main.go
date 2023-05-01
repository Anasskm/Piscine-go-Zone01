package main

import (
	"fmt"
	"os"
)

func main() {
	a := os.Args[1:]
	if len(a) == 0 {
		fmt.Println("File name missing")
		os.Exit(0)
	}
	if len(a) > 1 {
		fmt.Println("Too many arguments")
		os.Exit(0)
	}
	file, _ := os.Open(a[0])
	b := make([]byte, 25)
	file.Read(b)
	fmt.Println(string(b))
	file.Close()
}
