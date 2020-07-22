package main

import (
	"fmt"
	"strings"
)

func main() {
	const row1 = "VRCGVVRVCGGCCGVRGCVCGCGV"
	const row2 = "VRCCCGCRRGVCGCRVVCVGCGCV"
	findplants := "Charlie"

	var kids map[string]int
	var seeds map[string]string
	kids = map[string]int{
		"Alice":   1,
		"Bob":     2,
		"Charlie": 3,
		"David":   4,
		"Eve":     5,
		"Fred":    6,
		"Ginny":   7,
		"Harriet": 8,
		"Ileana":  9,
		"Joseph":  10,
		"Kincaid": 11,
		"Larry":   12,
	}

	seeds = map[string]string{
		"C": "clover",
		"G": "grass",
		"R": "radishes",
		"V": "violets",
	}

	kididx := kids[findplants]
	kidplants := row1[(kididx*2)-2:kididx*2] + row2[(kididx*2)-2:kididx*2]

	fmt.Printf("%v is growing: ", findplants)
	for i, v := range kidplants {
		if i == 0 {
			fmt.Printf("%v", strings.Title(seeds[string(v)]))
		} else {
			fmt.Printf("%v", seeds[string(v)])
		}
		if i != 3 {
			fmt.Printf(", ")
		}
	}
	fmt.Printf("\n")

}
