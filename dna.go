package main

import (
	"fmt"
)

func main() {
	const DNA1 = "GAGCCTACTAACGGGAT"
	const DNA2 = "CATCGTAATGACGGCCT"
	str1 := DNA1
	str2 := DNA2
	var diffs int

	for i := 0; i < len(str1); i++ {
		c := string(str1[i])
		d := string(str2[i])
		if c != d {
			//			fmt.Printf("c = %v d = %v\n", c, d)
			diffs++
		}
	}
	fmt.Printf("The Hamming Distance is %v\n", diffs)
}
