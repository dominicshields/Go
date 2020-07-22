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
	target := "listeners"
	candidates := []string{"google", "inlets", "enlisters", "banana", "inlet", "enlistersw"}
	var targetcnt []targetcnts
	var candidatecnt []targetcnts
	ok := 0

	targetcnt = createletterscore(target)

	for _, z := range candidates {
		candidatecnt = createletterscore(z)
		if len(z) == len(targetcnt) {
			for d := range candidatecnt {
				if candidatecnt[d].letter == targetcnt[d].letter && candidatecnt[d].occurs == targetcnt[d].occurs {
					ok = 1
				} else {
					ok = 0
					break
				}
			}
			if ok == 1 {
				fmt.Printf("%v and %v are anagrams\n", target, z)
			}
		}
		candidatecnt = candidatecnt[:0]
	}
}

func createletterscore(target string) []targetcnts {
	var targetcnt []targetcnts
	for _, r := range target {
		c := strings.ToLower(string(r))
		targetcnt = append(targetcnt, targetcnts{c, 1})
	}
	targetcnt = fixarray(targetcnt)
	return targetcnt
}

func fixarray(targetcnt []targetcnts) []targetcnts {
	sort.Slice(targetcnt, func(i, j int) bool { return targetcnt[i].letter < targetcnt[j].letter })
	for m := 0; m < len(targetcnt)-1; m++ {
		if targetcnt[m].letter == targetcnt[m+1].letter {
			targetcnt[m].occurs++
			targetcnt[m+1].letter = ""
			targetcnt[m+1].occurs = 0
			sort.Slice(targetcnt, func(i, j int) bool { return targetcnt[i].letter < targetcnt[j].letter })
		}
	}
	return targetcnt
}
