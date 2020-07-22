package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	const base = 5
	const numberstr = "82"
	result := 0
	numlength := len(numberstr)
	n := numlength -1
	for i := 0; i < numlength; i++ {
		x, _ := strconv.Atoi(numberstr[i : i+1])
		result += x * int(math.Pow(float64(base), float64(n)))
		n--
	}
	fmt.Println(result)
}