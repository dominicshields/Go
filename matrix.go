package main

import (
	"fmt"
	"strings"
)

func main() {
	var matrix string = "9 8 7\n5 3 2\n6 6 7"
	fmt.Println("ROWS")
	fmt.Println(strings.Replace(matrix, " ", ",", -1))
	fmt.Println("COLUMNS")
	x := 0
	for i := 0; i <= 2; i++ {
		fmt.Printf("%s,%s,%s\n", matrix[x:x+1], matrix[x+6:x+7], matrix[x+12:x+13]) 
		x += 2
	}
}