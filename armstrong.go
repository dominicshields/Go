package main

import (
	"fmt"
	"math"
	"strconv"
)

// http://mathworld.wolfram.com/NarcissisticNumber.html
//153, 370, 371, 407, 1634, 8208, 9474 54748, 92727, 93084, 548834, 1741725, 4210818, 9800817, 9926315,
//24678050, 24678051, 88593477, 146511208, 472335975, 534494836, 912985153, 4679307774

func main() {
	const testnum = 24678051
	sumofpowers := 0.0
	num := 0
	numfloat64 := 0.0
	testchar := strconv.Itoa(testnum)
	numlength := float64(len(testchar))

	for _, y := range testchar {
		num, _ = strconv.Atoi(string(y))
		numfloat64 = float64(num)
		sumofpowers += math.Pow(numfloat64, numlength)
	}
	fmt.Printf("Testchar = %v Armstrong Number = %.0f\n", testchar, sumofpowers)
}
