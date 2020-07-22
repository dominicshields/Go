package main

import (
	"fmt"
)

func main() {
	var compass map[int]string
	compass = map[int]string{
		0:  "North",
		1:  "East",
		-3: "East",
		2:  "South",
		-2: "South",
		-1: "West",
		3:  "West",
	}
	const instructions = "RAALARALARALLA"
	var currentdir string
	x := 7
	y := 3
	orientation := 0

	for _, v := range instructions {
		if string(v) == "R" {
			orientation++
			currentdir = "R"
		} else if string(v) == "L" {
			orientation--
			currentdir = "L"
		} else {
			if currentdir == "R" {
				x++
			} else {
				y++
			}
		}
	}

	if orientation > 3 || orientation < -3 {
		orientation = orientation % 4
	}

	fmt.Printf("Current coordinates are %v,%v facing %v\n", x, y, compass[orientation])
}
