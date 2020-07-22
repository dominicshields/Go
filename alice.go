package main

import (
	"fmt"
)

func main() {
	const ALICE = "Alice"

	//When X is a name or "you".
	//If the given name is "Alice", the result should be "One for Alice, one for me." If no name is given, the result should be "One for you, one for me."

	X := "cs"
	if X == ALICE {
		fmt.Printf("One for %s, one for me.\n", X)
	} else {
		fmt.Printf("One for you, one for me.\n")
	}

}