package main

import (
	"fmt"
	"math"
)

func main() {
	const LEFTBRACKET = "["
	const RIGHTBRACKET = "]"
	const LEFTBRACE = "{"
	const RIGHTBRACE = "}"
	const LEFTPAREN = "("
	const RIGHTPAREN = ")"
	bracket := 0
	brace := 0
	paren := 0

	phrase := "Given a string containing [brackets []]], braces { and parentheses ()))), verify"

	for _, v := range phrase {
		switch string(v) {
		case LEFTBRACKET:
			bracket++
		case RIGHTBRACKET:
			bracket--
		case LEFTBRACE:
			brace++
		case RIGHTBRACE:
			brace--
		case LEFTPAREN:
			paren++
		case RIGHTPAREN:
			paren--
		}
	}

	if bracket == 0 {
		fmt.Printf("The square brackets balance\n")
	} else if bracket < 0 {
		fmt.Printf("There is an excess of %v right square brackets\n", math.Abs(float64(bracket)))
	} else {
		fmt.Printf("There is an excess of %v left square brackets\n", math.Abs(float64(bracket)))
	}

	if brace == 0 {
		fmt.Printf("The curly brackets balance\n")
	} else if brace < 0 {
		fmt.Printf("There is an excess of %v right curly brackets\n", math.Abs(float64(brace)))
	} else {
		fmt.Printf("There is an excess of %v left curly brackets\n",math.Abs(float64(brace)))
	}

	if paren == 0 {
		fmt.Printf("The round brackets balance\n")
	} else if paren < 0 {
		fmt.Printf("There is an excess of %v right round brackets\n",math.Abs(float64(paren)))
	} else {
		fmt.Printf("There is an excess of %v left round brackets\n",math.Abs(float64(paren)))
	}

	//	fmt.Printf("bracket = %v, brace = %v, paren = %v\n", bracket, brace, paren)
}
