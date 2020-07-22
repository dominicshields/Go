package main

import (
	"fmt"
	"math"
)

func main() {
	const numtocheck = 28
	fmt.Printf("%v is %v\n", numtocheck, classify(numtocheck))
}
func classify(numtocheck int) (retval string) {
	var aliquot int
	for y := 1; y < numtocheck; y++ {
		if math.Mod(float64(numtocheck), float64(y)) == 0 {
			aliquot += y
		}
	}
	if aliquot < numtocheck {
		retval = "Deficient"
	} else if aliquot == numtocheck {
		retval = "Perfect"
	} else {
		retval = "Abundant"
	}
	return retval
}
