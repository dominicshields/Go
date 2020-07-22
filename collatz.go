package main

import (
	"fmt"
)

func main() {
	count := 0

	for testnum := 12; testnum > 1; count++ {
		if testnum%2 == 0 {
			testnum = testnum / 2
		} else {
			testnum *= 3
			testnum++
		}
		fmt.Println(testnum)
	}
	fmt.Printf("Number of steps = %v\n", count)
}