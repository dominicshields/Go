package main

import (
	"fmt"
)

type Triangle struct {
	side1 int
	side2 int
	side3 int
}

func main() {
	const side1len = 7
	const side2len = 5
	const side3len = 7
	const EQUILATERAL = "Equilateral"
	const ISOCELES = "Isoceles"
	const SCALENE = "Scalene"
	t := Triangle{}
	t.side1 = side1len
	t.side2 = side2len
	t.side3 = side3len

	if t.side1 == t.side2 && t.side1 == t.side3 {
		fmt.Printf("Triangle is %s\n", EQUILATERAL)
	} else if t.side1 == t.side2 && t.side1 != t.side3 || t.side1 != t.side2 && t.side1 == t.side3 {
		fmt.Printf("Triangle is %s\n", ISOCELES)
	} else if t.side1 != t.side2 && t.side1 != t.side3 && t.side2 != t.side3 {
		fmt.Printf("Triangle is %s\n", SCALENE)
	}

}
