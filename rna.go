package main

import (
	"fmt"
)

func main() {
	var RNA []string
	var codes = map[string]string{
		"A": "U",
		"C": "G",
		"G": "C",
		"T": "A",
	}
	const DNA = "ACGTTTGAA"

	for _, x := range DNA {
		RNA = append(RNA, codes[string(x)])
	}

	fmt.Println(RNA)
}
