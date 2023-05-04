package main

import "github.com/01-edu/z01"

type Door struct {
	state string
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func CloseDoor(Door *Door) {
	PrintStr("Door Closing...")
	Door.state = "CLOSE"
}

func IsDoorOpen(Door *Door) bool {
	PrintStr("is the Door opened ?")
	if Door.state == "OPEN" {
		return true
	} else {
		return false
	}
}

func IsDoorClose(Door *Door) bool {
	PrintStr("is the Door closed ?")
	if Door.state == "CLOSE" {
		return true
	} else {
		return false
	}
}

func OpenDoor(Door *Door) {
	PrintStr("Door Opening...")
	Door.state = "OPEN"
}

func main() {
	door := &Door{}

	CloseDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(door) {
		CloseDoor(door)
	}
}
