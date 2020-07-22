package main

import (
	"fmt"
	"math"
)

func main() {
	const year = 2004
	if math.Mod(year, 4) == 0 {
		if math.Mod(year, 100) == 0 {
			if math.Mod(year, 400) == 0 {
				fmt.Printf("Year %d is a leap year\n", year)
			} else {
				fmt.Printf("Year %d is not a leap year\n", year)

			}
		} else {
			fmt.Printf("Year %d is a leap year\n", year)

		}
	} else {
		fmt.Printf("Year %d is not a leap year\n", year)
	}

}