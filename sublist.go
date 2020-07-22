package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	B := []int{1, 2, 33, 4, 5}
	A := []int{2, 33, 4}
	var comparisonA string
	var comparisonB string

	for x, _ := range B {
		comparisonB = comparisonB + strconv.Itoa(B[x])
	}

	for y, _ := range A {
		comparisonA = comparisonA + strconv.Itoa(A[y])
	}

	if comparisonA == comparisonB {
		fmt.Printf("A: %v is equal to B: %v\n", A, B)
	} else if strings.Contains(comparisonB, comparisonA) {
		fmt.Printf("A: %v is a sublist of B: %v\n", A, B)
	} else if strings.Contains(comparisonA, comparisonB) {
		fmt.Printf("A: %v is a superlist of B: %v\n", A, B)
	} else {
		fmt.Printf("A: %v is not a superlist of, sublist of or equal to B: %v\n", A, B)
	}

}

