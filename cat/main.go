package main

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func PrintResult(str string) {
	for _, val := range str {
		z01.PrintRune(rune(val))
	}
}

func Read(file string) string {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return "error"
	}
	return string(content)
}

func main() {
	args := os.Args[1:]
	// abc := "1"
	finish := false
	for _, file := range args {
		if _, err := os.Stat(file); err != nil {
			PrintResult("ERROR: open " + file + ": no such file or directory\n")
			os.Exit(1)
		}
		PrintResult(Read(file))
		finish = true
	}
	if !finish {
		reader := io.TeeReader(os.Stdin, os.Stdout)
		ioutil.ReadAll(reader)
		os.Stdin.Close()
		os.Stdout.Close()
	}
}
