package main

import (
	"fmt"
)

func main() {

	// Slice passed by value
	var slice1 []int = []int{1, 2, 3}
	fmt.Printf("slice1 before: %v\n", slice1)
	sliceByVal(slice1)
	fmt.Printf("slice1 after:  %v\n", slice1)

	// Slice passed by reference
	var slice2 []int = []int{1, 2, 3}
	fmt.Printf("slice2 before: %v\n", slice2)
	sliceByRef(&slice2)
	fmt.Printf("slice2 after:  %v\n", slice2)

	// Slice backed by array passed by value
	var array1 [6]int = [6]int{0, 1, 2, 3, 5, 6}
	slice3 := array1[1:4]
	fmt.Printf("array1 before: %v\n", array1)
	sliceByVal(slice3)
	fmt.Printf("array1 after:  %v\n", array1)

	// Slice backed by array passed by reference
	var array2 [6]int = [6]int{0, 1, 2, 3, 5, 6}
	slice4 := array2[1:4]
	fmt.Printf("array2 before: %v\n", array2)
	sliceByRef(&slice4)
	fmt.Printf("array2 after:  %v\n", array2)

	//Modify the contents of a slice passed by value	
	fmt.Printf("slice1 before: %v\n", slice1)
	sliceByVal2(slice1)
	fmt.Printf("slice1 after:  %v\n", slice1)

	// Slice passed by reference
	fmt.Printf("slice2 before: %v\n", slice2)
	sliceByRef2(&slice2)
	fmt.Printf("slice2 after:  %v\n", slice2)


}

func sliceByVal(slice []int) {
	slice = append(slice, 4)
}

func sliceByRef(slice *[]int) {
	*slice = append(*slice, 4)
}

func sliceByVal2(slice []int) {
	slice[0] = 42	
}

func sliceByRef2(slice *[]int) {
	(*slice)[0] = 1337
}

