package main

import (
	"fmt"
)

type scores struct {
	letters string
	value   int
}

func main() {
	dnastrand := "AGCTTTTCATTCTGACTGCAACGGGCAATATGTCTCTGTGTGGATTAAAAAAAGAGTGTCTGATAGCAGC"
	flag := 0
	score := []scores{}
	score = append(score, scores{"A", 0})
	score = append(score, scores{"C", 0})
	score = append(score, scores{"G", 0})
	score = append(score, scores{"T", 0})

	for i := 0; i < len(dnastrand); i++ {
		c := string(dnastrand[i])
		flag = 0
		for k, _ := range score {

			if score[k].letters == c {
				score[k].value++
				flag = 1
			}
		}
		if flag != 1 {
			fmt.Printf("Invalid characters found %v in %v\n", c, dnastrand)
			break
		}
	}
	if flag == 1 {
		fmt.Printf("Histogram %v\n", score)
	}
}
