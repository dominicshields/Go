package main

import (
	"fmt"
)

var letter map[string]int

const key = "somekey"
const plaintext = "thequickbrownfoxjumpedoverthelazydog"

func main() {
	letter = map[string]int{
		"a": 0,
		"b": 1,
		"c": 2,
		"d": 3,
		"e": 4,
		"f": 5,
		"g": 6,
		"h": 7,
		"i": 8,
		"j": 9,
		"k": 10,
		"l": 11,
		"m": 12,
		"n": 13,
		"o": 14,
		"p": 15,
		"q": 16,
		"r": 17,
		"s": 18,
		"t": 19,
		"u": 20,
		"v": 21,
		"w": 22,
		"x": 23,
		"y": 24,
		"z": 25,
	}

	caesar := improvedcaesar(13)
	fmt.Printf("Caesar = %v\n", caesar)

	ciphertext := encode(plaintext)
	fmt.Printf("Vigenere encoded = %v\n", ciphertext)

	deciphertext := decode(ciphertext)
	fmt.Printf("Vigenere decoded = %v\n", deciphertext)

}

func encode(plaintext string) string {
	var ciphertext string
	var x string
	for i, v := range plaintext {
		if i < len(key) {
			x = key[i : i+1]
		} else {
			mult := int(i / len(key))
			x = key[i-len(key)*mult : i-(len(key)*mult)+1]
		}
		shift := letter[x]
		topletter := string('z' - shift)
		reverseshift := 26 - shift
		if string(v) <= topletter {
			ciphertext += string(int(v) + shift)
		} else {
			ciphertext += string(int(v) - reverseshift)
		}
	}
	return ciphertext
}

func decode(ciphertext string) string {
	var plaintext string
	var x string
	for i, v := range ciphertext {
		if i < len(key) {
			x = key[i : i+1]
		} else {
			mult := int(i / len(key))
			x = key[i-len(key)*mult : i-(len(key)*mult)+1]
		}
		shift := letter[x]
		bottomletter := string('a' + shift)
		reverseshift := 26 - shift
		if string(v) >= bottomletter {
			plaintext += string(int(v) - shift)
		} else {
			plaintext += string(int(v) + reverseshift)
		}
	}
	return plaintext
}

func improvedcaesar(shift int) string {
	const alphabet = 26
	reverseshift := alphabet - shift
	topletter := string('z' - shift)
	const plaintext = "thequickbrownfoxjumpsoverthelazydog"
	var ciphertext string

	for _, v := range plaintext {
		if string(v) <= topletter {
			ciphertext += string(int(v) + shift)
		} else {
			ciphertext += string(int(v) - reverseshift)
		}
	}
	return ciphertext
}
