package main

import (
	"fmt"
)

func main() {
	const N = 64
	var grains uint64 = 1
	var squareval uint64 = 1

	for x := 2; x <= N; x++ {
		squareval = squareval * 2
		grains += squareval
		fmt.Printf("Square %d contains %v grains of wheat\n", x, squareval)

	}

	fmt.Printf("Number of grains = %v\n", grains)
	//18446744073709551615
}
