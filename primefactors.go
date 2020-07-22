package main

import (
	"fmt"
)

func main() {
	const testnum = 24
	var factors []int
	var factor int
	total := 1
	factor = testnum

	for i := 2; i < testnum/2+1; i++ {
		for factor%i == 0 {
			if factor%i == 0 {
				factor = factor / i
				if i != 1 {
					factors = append(factors, i)
				}
			} else {
				break
			}
		}
	}

	for _, y := range factors {
		total *= y
	}

	fmt.Printf("Number tested = %v, Prime Factors = %v, Product of factors = %v\n", testnum, factors, total)
}
