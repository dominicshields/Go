package main

import (
	"fmt"
	"sort"
	"strings"
)

type scores struct {
	letters string
	value   int
}

type newscores struct {
	letters string
	value   int
}

func main() {
	score := []scores{}
	newscore := []newscores{}
	score = append(score, scores{"AEIOULNRST", 1})
	score = append(score, scores{"DG", 2})
	score = append(score, scores{"BCMP", 3})
	score = append(score, scores{"FHVWY", 4})
	score = append(score, scores{"K", 5})
	score = append(score, scores{"JX", 8})
	score = append(score, scores{"QZ", 10})

	for k, _ := range score {
		for _, c := range score[k].letters {
			newscore = append(newscore, newscores{strings.ToLower(string(c)), score[k].value})
		}
	}

	sort.Slice(newscore, func(i, j int) bool { return newscore[i].letters < newscore[j].letters })

	for k, _ := range newscore {
		fmt.Printf("%v is worth %v points\n", newscore[k].letters, newscore[k].value)
	}

}
