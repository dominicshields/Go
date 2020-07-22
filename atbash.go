package main

import (
	"fmt"
)

const alphamax = 26

var cyphertable [alphamax][2]string

func main() {
	const plaintext = "thequickbrownfoxjumpsoverthelazydog"
	setup()

	cyphertext := encode(plaintext)
	fmt.Println(cyphertext)

	recoveredplaintext := decode(cyphertext)
	fmt.Println(recoveredplaintext)
}

func setup() {
	for i := 0; i < alphamax; i++ {
		cyphertable[i][0] = string('a' + i)
	}

	for i := alphamax - 1; i >= 0; i-- {
		cyphertable[i][1] = string('z' - i)
	}
	fmt.Println(cyphertable)
}

func encode(plaintext string) string {
	cnt := 0
	var cyphertext string
	for x, y := range plaintext {
		for x = range cyphertable {
			if cyphertable[x][0] == string(y) {
				if cnt == 5 {
					cyphertext += " "
					cnt = 0
				}
				cyphertext += cyphertable[x][1]
				cnt++
			}
		}
	}
	return cyphertext
}

func decode(cyphertext string) string {
	var plaintext string
	for x, y := range cyphertext {
		for x = range cyphertable {
			if cyphertable[x][1] == string(y) {
				plaintext += cyphertable[x][0]
			}
		}
	}
	return plaintext
}
