package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	const alphamax = 26
	const dot = "."
	diamond := "L"
	var alphabet [alphamax]string
	var output []string
	letteridx := 0

	if len(os.Args) > 1 {
		diamond = os.Args[1]
	}

	for i := 0; i < alphamax; i++ {
		alphabet[i] = string('A' + i)
	}

	for x := 0; x < alphamax; x++ {
		if alphabet[x] == diamond {
			letteridx = x
		}
	}

	for w := 0; w <= letteridx; w++ {
		if w == 0 {
			output = append(output, strings.Repeat(dot, letteridx-w)+alphabet[w]+strings.Repeat(dot, letteridx))
		} else {
			output = append(output, strings.Repeat(dot, letteridx-w)+alphabet[w]+strings.Repeat(dot, w*2-1)+alphabet[w]+strings.Repeat(dot, letteridx-w))
		}
	}

	for y := letteridx - 1; y >= 0; y-- {
		if y == 0 {
			output = append(output, strings.Repeat(dot, letteridx-y)+alphabet[y]+strings.Repeat(dot, letteridx))
		} else {
			output = append(output, strings.Repeat(dot, letteridx-y)+alphabet[y]+strings.Repeat(dot, y*2-1)+alphabet[y]+strings.Repeat(dot, letteridx-y))
		}
	}

	for x := range output {
		fmt.Printf("%v\n", output[x])
	}

}
