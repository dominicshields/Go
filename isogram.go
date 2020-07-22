package main

import (
	"fmt"
	//	"strings"
)

func main() {
	var cnt int
	scabbleword := "AMBIDEXTROUSLY"

	for i := 0; i < len(scabbleword); i++ {
		c := string(scabbleword[i])
		cnt = 0
		for x := 0; x < len(scabbleword); x++ {
			d := string(scabbleword[x])
			if c == d {
				cnt++
			}
		}
	}
	if cnt < 2 {
		fmt.Printf("%s is an Isogram\n", scabbleword)
	} else {
		fmt.Printf("%s is NOT an Isogram\n", scabbleword)
	}
}
