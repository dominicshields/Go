package main

import (
	"fmt"
	"strings"
)

type alphabet struct {
	letter string
	occurs int
}

func main() {
	//	const sentence = "The quick brown fox jumps over the lazy dog"
	const sentence = "Squdgy fez, blank jimp crwth vox"
	var alpha []alphabet
	flag := 0

	for i := 1; i < 27; i++ {
		alpha = append(alpha, alphabet{string('a' - 1 + i), 0})
	}

	for _, r := range sentence {
		c := strings.ToLower(string(r))
		for k, _ := range alpha {
			if strings.Contains(alpha[k].letter, c) {
				alpha[k].occurs++
			}
		}
	}

	for k, _ := range alpha {
		if alpha[k].occurs == 0 {
			fmt.Printf("The Letter %v does not occur\n", alpha[k].letter)
			flag = 1
		}
	}

	if flag == 0 {
		fmt.Printf("The following sentence is a Pangram: %v\n", sentence)
	}
}
