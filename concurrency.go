package main

import (
	"fmt"
	"strings"
	"sync"
)

type scores struct {
	letter string
	value  int
}

var wg sync.WaitGroup

func main() {
	var textarr [2]string
	textarr[0] = "AMBIDEXTROUSLY"
	textarr[1] = "ANTIDISESTABLISHMENTARIANISM"
	for z := 0; z < len(textarr); z++ {
		wg.Add(1)
		go concurrent(textarr[z])
	}
	wg.Wait()
}
func concurrent(text string) {
	score := []scores{}
	for x := 1; x <= 26; x++ {
		score = append(score, scores{toChar(x), 0})
	}

	for i := 0; i < len(text); i++ {
		c := string(text[i])

		for k, _ := range score {
			if strings.Contains(score[k].letter, c) {
				score[k].value += 1
			}
		}
	}

	for k, _ := range score {
		fmt.Printf("%s : Letter %v appears %v\n", text, score[k].letter, score[k].value)
	}
	wg.Done()
}

func toChar(i int) string {
	return string('A' - 1 + i)
}
