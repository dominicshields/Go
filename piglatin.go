package main

import (
	"fmt"
)

func main() {
	const suffix = "ay"
	const testword = "dominic"
	vowels := []string{"a", "e", "i", "o", "u"}
	var latin string
	vowelfound := 0

	prefix := testword[0:1]
	for _, y := range vowels {
		if prefix == y {
			vowelfound = 1
		}
	}
	if vowelfound == 1 {
		latin = testword + suffix
	} else {
		latin = testword[1:] + prefix + suffix
	}

	fmt.Printf("Latin word = %v\n", latin)
}
