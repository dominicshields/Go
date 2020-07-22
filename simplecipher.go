package main

import (
	"fmt"
)

func main() {
	const shift = 13
	const reverseshift = 26 - shift
	const topletter = string('z' - shift)
	const plaintext = "thequickbrownfoxjumpsoverthelazydog"
	var ciphertext string

	for _, v := range plaintext {
		if string(v) <= topletter {
			ciphertext += string(v + shift)
		} else {
			ciphertext += string(v - reverseshift)
		}
	}

	fmt.Println(ciphertext)

}
