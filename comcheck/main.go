package main

import (
	"fmt"
	"os"
)

func main() {
	a := os.Args[1:]
	for _, i := range a {
		if i == "galaxy" || i == "01" || i == "galaxy 01" {
			fmt.Println("Alert!!!")
			os.Exit(0)
		}
	}
	os.Exit(0)
}
