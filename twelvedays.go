package main

import (
	"fmt"
)

func main() {
	const fixed1 = "On the "
	const fixed2 = " day of Christmas my true love gave to me, "
	var day = [12]string{"first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth"}
	var gift = [12]string{"a Partridge in a Pear Tree.", "two Turtle Doves", "three French Hens", "four Calling Birds", "five Gold Rings", "six Geese-a-Laying", "seven Swans-a-Swimming", "eight Maids-a-Milking", "nine Ladies Dancing", "ten Lords-a-Leaping", "eleven Pipers Piping", "twelve Drummers Drumming"}
	var w int

	for i := 0; i < 12; i++ {
		fmt.Printf("%v%v%v%v", fixed1, day[i], fixed2, gift[i])

		for w = i - 1; w >= 0; w-- {
			if w == 0 {
				fmt.Printf(" and ")
			} else {
				fmt.Printf(", ")
			}
			fmt.Printf("%v", gift[w])
		}
		fmt.Printf("\n")
	}

}