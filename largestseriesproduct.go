package main

import (
	"fmt"
	"strconv"
)

func main() {
	const length = 6
	const num = "73167176531330624919225119674426574742355349194934"
	var numhigh string
	var numeval string
	var multeval int
	var multhigh int

	for t := 0; t < len(num)-length; t++ {
		numeval = num[t : t+length]
		multeval = 0
		for x := 0; x < len(numeval); x++ {
			i, _ := strconv.Atoi(numeval[x : x+1])
			if multeval == 0 {
				multeval = i
			} else {
				multeval *= i
			}
			//			fmt.Printf("multeval = %v, i = %v\n", multeval, i)
		}
		if multeval > multhigh {
			multhigh = multeval
			numhigh = numeval
		}
	}
	fmt.Printf("Highest %v adjacent numbers = %v\n", length, numhigh)
	fmt.Printf("Highest product of %v adjacent numbers = %v\n", length, multhigh)
}
