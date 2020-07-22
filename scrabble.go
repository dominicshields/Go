package main

import (
	"fmt"
	"strings"
)

type scores struct {
	letters string
	value   int
}

func main() {
	scabbleword := "AMBIDEXTROUSLY"
	score := []scores{}
	var wordscore int
	score = append(score, scores{"AEIOULNRST", 1})
	score = append(score, scores{"DG", 2})
	score = append(score, scores{"BCMP", 3})
	score = append(score, scores{"FHVWY", 4})
	score = append(score, scores{"K", 5})
	score = append(score, scores{"JX", 8})
	score = append(score, scores{"QZ", 10})

	for i := 0; i < len(scabbleword); i++ {
		c := string(scabbleword[i])

		for k, _ := range score {
			if strings.Contains(score[k].letters, c) {
				wordscore += score[k].value
			}
		}

	}
	fmt.Printf("The Scrabble Word Score for %v is %v\n", scabbleword, wordscore)
}