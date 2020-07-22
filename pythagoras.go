package main

import (
	"fmt"
)

func main() {
	limit := 1000
	for a := 0; a <= limit/2; a++ {

		for b := 0; b <= limit/2; b++ {

			for c := 0; c <= limit/2; c++ {

				if (a*a)+(b*b) == c*c && a+b+c == limit && a > 0 && b > 0 {
					fmt.Printf("a = %v, b = %v, c = %v\n", a, b, c)
				}
			}
		}
	}
}
