package main

import (
	"fmt"
)

type decromans struct {
	dec   int
	roman string
}

func main() {
	var decroman = []decromans{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	//const DECYEAR = 1950 // MCML
	//const DECYEAR = 2018 // MMXVIII
	const DECYEAR = 2009 // MMIX
	var numerals string
	y := 0
	x := 0

	for y = DECYEAR; y > 0; y -= decroman[x].dec {
		for x = 0; x < len(decroman); x++ {
			if y >= decroman[x].dec {
				numerals += decroman[x].roman
				break
			}
		}
	}
	fmt.Printf("%v = %v\n", DECYEAR, numerals)
}
