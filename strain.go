package main

import (
	"fmt"
)

func main() {
	const predicate = "Even"
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Printf("All Nums: %v\n", nums)
	if predicate == "Even" {
		fmt.Printf("Keep Nums: %v\n", keep(nums, predicate))
	} else {
		fmt.Printf("Discard Nums: %v\n", discard(nums, predicate))
	}

}
func discard(nums []int, predicate string) []int {
	var retnums []int
	for _, y := range nums {
		if predicate != "Even" {
			if y%2 != 0 {
				retnums = append(retnums, y)
			}
		}
	}

	return retnums
}
func keep(nums []int, predicate string) []int {
	var retnums []int
	for _, y := range nums {
		if y%2 == 0 {
			retnums = append(retnums, y)
		}
	}

	return retnums
}
