package main

import (
	"fmt"
	"sort"
)

var tot int
var multiples []int

func main() {
	limit := 40
	const numa = 3
	const numb = 5
	var a int
	var b int

	for i := 1; i <= limit/numa; i++ {
		a = numa * i
		check(a, limit)
		b = numb * i
		check(b, limit)
	}
	sort.Ints(multiples)
	fmt.Println("Multiples", multiples)
	fmt.Printf("Sum of multiples = %v\n", tot)
}

func check(numin int, limit int) {
	var flag int
	if numin < limit {
		flag = 0
		for i := range multiples {
			if numin == multiples[i] {
				flag = 1
			}
		}
		if flag == 0 {
			tot = tot + numin
			multiples = append(multiples, numin)
		}
	}
}
