package main

import (
	"fmt"
)

const coinmax = 7

func main() {
	const changereqd = 113
	var coins = [coinmax]int{100, 50, 20, 10, 5, 2, 1}
	var change []int

	changevar := changereqd
	for i := 0; i < coinmax; i++ {
		if coins[i] <= changevar {
			change = append(change, coins[i])
			changevar -= coins[i]
			if changevar >= coins[i] {
				i--
			}
		}
	}

	fmt.Println(change)
}