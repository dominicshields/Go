package main

import (
	"fmt"
	"sort"
	"strings"
)

type wordcnts struct {
	word   string
	occurs int
}

func main() {
	const sentence = "word A sample word sentence word to test on word"
	var words []string
	var wordcnt []wordcnts

	words = strings.Split(sentence, " ")
	for _, y := range words {
		wordcnt = append(wordcnt, wordcnts{strings.ToLower(y), 1})
	}

	wordcnt = fixarray(wordcnt)

	for r := range wordcnt {
		if wordcnt[r].occurs > 0 {
			fmt.Printf("%v: %v\n", wordcnt[r].word, wordcnt[r].occurs)
		}
	}
}

func fixarray(targetcnt []wordcnts) []wordcnts {
	sort.Slice(targetcnt, func(i, j int) bool { return targetcnt[i].word < targetcnt[j].word })
	for m := 0; m < len(targetcnt)-1; m++ {
		if targetcnt[m].word == targetcnt[m+1].word {
			targetcnt[m].occurs++
			targetcnt[m+1].word = ""
			targetcnt[m+1].occurs = 0
			sort.Slice(targetcnt, func(i, j int) bool { return targetcnt[i].word < targetcnt[j].word })
		}
	}
	return targetcnt
}
