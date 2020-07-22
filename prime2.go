package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	const LIMIT = 10000
	var allnums [LIMIT]int
	var cnt int

	file, err := os.Create("d:\temp\result.txt")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

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
			fmt.Fprintf(file, "%d\n", i)
			cnt++
		}
	}
	fmt.Println("Total number of primes = ", cnt)

}
