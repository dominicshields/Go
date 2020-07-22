package main

import (
	"fmt"
)

func main() {
	var reverse string
	scabbleword := "AMBIDEXTROUSLY"

	for x := len(scabbleword) - 1; x >= 0; x-- {
		d := string(scabbleword[x])
		reverse = reverse + d
	}

	fmt.Printf("%s reversed is %s\n", scabbleword, reverse)
}
