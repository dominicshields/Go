package main

import (
	"fmt"
)

func main() {
	const PLING3 = "Pling"
	const PLANG5 = "Plang"
	const PLONG7 = "Plong"
	//	var flag bool
	num := 105
	flag := false
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			switch i {
			case 3:
				fmt.Printf("%v", PLING3)
				flag = true
			case 5:
				fmt.Printf("%v", PLANG5)
				flag = true
			case 7:
				fmt.Printf("%v", PLONG7)
				flag = true
			}
		}
	}
	if !flag {
		fmt.Printf("%v", num)
	}
}