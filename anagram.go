package main

import (
	"fmt"
	"sort"
	"strings"
)

type targetcnts struct {
	letter string
	occurs int
}

func main() {
	target := "listtenerse"
	candidates := []string{"enlists", "google", "inlets", "banana", "inlett"}
	i := 0
	var targetcnt []targetcnts

	for _, r := range target {
		c := strings.ToLower(string(r))
		targetcnt = append(targetcnt, targetcnts{c, 1})
	}

	sort.Slice(targetcnt, func(i, j int) bool { return targetcnt[i].letter < targetcnt[j].letter })

	for m := 0; m < len(targetcnt)-1; m++ {
		if targetcnt[m].letter == targetcnt[m+1].letter {
			targetcnt[m].occurs++
			targetcnt[m+1].letter = ""
			targetcnt[m+1].occurs = 0
			sort.Slice(targetcnt, func(i, j int) bool { return targetcnt[i].letter < targetcnt[j].letter })
		}
	}

	fmt.Printf("targetcnt = %v\n", targetcnt)

	for x := range candidates {
		for _, y := range candidates[x] {
			if strings.ContainsRune(target, y) {
				i = 0
			} else {
				i = 1
				break
			}
		}
		if i == 0 && len(target) == len(candidates[x]) {
			fmt.Printf("%v is an anagram of %v\n", candidates[x], target)
		}
	}
}
