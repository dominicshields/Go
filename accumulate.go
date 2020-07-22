package main

import (
	"fmt"
)

func main() {
	var array = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var squares []int

	for x := range array {
		//	fmt.Printf("%d\n", array[x]*array[x])
		squares = append(squares, array[x]*array[x])

	}
	fmt.Println(array)
	fmt.Println(squares)
}
