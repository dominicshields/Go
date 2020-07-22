package main

import (
	"fmt"
)

func main() {
	const LIMIT = 10000
	var allnums [LIMIT]int
	var cnt int

	for i := 2; i < LIMIT; i++ {
		allnums[i] = i
	}

	for i := 2; i < LIMIT; i++ {
		if allnums[i] != 0 {
			for j := i; i*j < LIMIT; j++ {

				allnums[i*j] = 0
			}
		}
	}

	for i := 2; i < LIMIT; i++ {
		if allnums[i] != 0 {
			fmt.Println(i)
			cnt++
		}
	}
	fmt.Println("Total number of primes = ", cnt)

}