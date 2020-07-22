package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {

	const phrase = "Portable Network Graphics"
	var buffer bytes.Buffer
	arr := strings.Split(phrase, " ")
	fmt.Printf("%q\n", strings.Split(phrase, " "))
	for _, each := range arr {
		//		fmt.Printf("[%s]", each[0:1])
		buffer.WriteString(each[0:1])
	}
	fmt.Println(buffer.String())
}