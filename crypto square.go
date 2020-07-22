package main

import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

func main() {

	const sentence = "If man was meant to stay on the ground, god would have given us roots."
	var square []string
	var cryptoslice []string
	var normsentence string
	var crypto string

	ccc := 0
	rrr := 0

	for _, x := range sentence {
		if unicode.IsLetter(x) {
			normsentence += strings.ToLower(string(x))
		}
	}

	s := math.Sqrt(float64(len(normsentence)))
	c := math.Ceil(s)
	r := math.Floor(s)

	cc := int(c)
	for i := 0; i < len(normsentence); i += int(c) {
		ccc = ccc + cc
		if ccc > len(normsentence) {
			ccc = len(normsentence)
		}
		square = append(square, normsentence[i:ccc])
	}

	for t := 0; t < int(c); t++ {
		for _, x := range square {
			if t < len(x) {
				crypto += x[t : t+1]
			}
		}
	}

	rr := int(r)
	for u := 0; u < len(crypto); u += int(r) {
		rrr += rr
		if rrr > len(crypto) {
			rrr = len(crypto)
		}
		cryptoslice = append(cryptoslice, crypto[u:rrr])
	}

	fmt.Println(square)
	fmt.Println(cryptoslice)
}
