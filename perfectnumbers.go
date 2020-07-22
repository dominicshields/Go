package main

import (
	"fmt"
	"math"
)

func main() {
	const numtocheck = 28
	var factors []int
	var aliquot int

	for y := 1; y < numtocheck; y++ {
		if math.Mod(float64(numtocheck), float64(y)) == 0 {
			factors = append(factors, y)
			aliquot += y
		}
	}
	fmt.Printf("Factors of %v = %v\n", numtocheck, factors)
	fmt.Printf("Aliquot Sum = %v\n", aliquot)
	if aliquot < numtocheck {
		fmt.Printf("%v is deficient\n", numtocheck)
	} else if aliquot == numtocheck {
		fmt.Printf("%v is perfect\n", numtocheck)
	} else {
		fmt.Printf("%v is abundant\n", numtocheck)
	}
}