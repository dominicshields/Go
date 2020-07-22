package main

import (
	"fmt"
	"strconv"
)

type decromans struct {
	dec   int
	roman string
}

var decroman = []decromans{
	{1000000000, "billion"},
	{1000000, "million"},
	{1000, "thousand"},
	{100, "hundred"},
	{90, "ninety"},
	{80, "eighty"},
	{70, "seventy"},
	{60, "sixty"},
	{50, "fifty"},
	{40, "forty"},
	{30, "thirty"},
	{20, "twenty"},
	{19, "nineteen"},
	{18, "eighteen"},
	{17, "seventeen"},
	{16, "sixteen"},
	{15, "fifteen"},
	{14, "fourteen"},
	{13, "thirteen"},
	{12, "twelve"},
	{11, "eleven"},
	{10, "ten"},
	{9, "nine"},
	{8, "eight"},
	{7, "seven"},
	{6, "six"},
	{5, "five"},
	{4, "four"},
	{3, "three"},
	{2, "two"},
	{1, "one"},
	{0, "zero"},
}

func main() {

	const LOWER = 0
	const UPPER = 2147483647 // int32 MAX
	const DECYEAR = 2103231111
	var numerals string
	x := 0
	d := 1

	if DECYEAR < LOWER || DECYEAR > UPPER {
		fmt.Printf("Number input %v out of range\n", DECYEAR)
		return
	}
	numberlen := len(strconv.Itoa(DECYEAR))
	y := DECYEAR

	for m := 1; m <= numberlen; m++ {
		d = 1
		x = lookup(y)

		if decroman[x].dec > 99 {
			d = int(y / decroman[x].dec)
			if d > 0 {

				if d > 100 {
					h := lookup(d / 100)
					numerals += decroman[h].roman + " "
				}

				e := lookup(d)
				numerals += decroman[e].roman + " "

				if d > 100 {
					g := lookup(d % 100)
					numerals += decroman[g].roman + " "
				}

				if d%100 > 20 {
					f := lookup(d % 10)
					numerals += decroman[f].roman + " "
				}
			}
		}

		numerals += decroman[x].roman + " "
		y -= d * decroman[x].dec
		if y < 1 {
			break
		}
	}
	fmt.Printf("%v = %v\n", DECYEAR, numerals)
}

func lookup(y int) int {
	x := 0
	for x = 0; x < len(decroman); x++ {
		if y >= decroman[x].dec {
			break
		}
	}
	return x
}
