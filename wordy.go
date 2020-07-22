package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	const computation = "What is 555 plus 13?"
	var result float64

	evaluate := computation[8 : len(computation)-1]
	i := strings.Index(evaluate, " ")
	firstnumstr := evaluate[:i]
	firstnum, _ := strconv.Atoi(firstnumstr)
	j := strings.LastIndex(evaluate, " ")
	lastnumstr := evaluate[j+1:]
	lastnum, _ := strconv.Atoi(lastnumstr)
	operation := strings.TrimSpace(evaluate[i:j])

	switch operation {
	case "plus":
		result = float64(firstnum) + float64(lastnum)
	case "minus":
		result = float64(firstnum) - float64(lastnum)
	case "multiplied by":
		result = float64(firstnum) * float64(lastnum)
	case "divided by":
		result = float64(firstnum) / float64(lastnum)
	}
	fmt.Printf("firstnum %v, lastnum %v, operation %v = %v\n", firstnum, lastnum, operation, result)
}
