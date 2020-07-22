package main

import (
	"fmt"
)

func main() {
	fmt.Println(hello())
}
func hello() string {
	dom := "Hello World!"
	return dom
}