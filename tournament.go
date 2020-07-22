package main

import (
	"fmt"
	"sort"
)

type results struct {
	home    string
	away    string
	outcome string
}

type leagues struct {
	team string
	MP   int
	W    int
	D    int
	L    int
	P    int
}

func main() {
	result := []results{}
	league := []leagues{}
	const WIN = "win"
	const DRAW = "draw"
	const LOSS = "loss"

	result = append(result, results{"Allegoric Alaskans", "Blithering Badgers", "win"})
	result = append(result, results{"Devastating Donkeys", "Courageous Californians", "draw"})
	result = append(result, results{"Devastating Donkeys", "Allegoric Alaskans", "win"})
	result = append(result, results{"Courageous Californians", "Blithering Badgers", "loss"})
	result = append(result, results{"Blithering Badgers", "Devastating Donkeys", "loss"})
	result = append(result, results{"Allegoric Alaskans", "Courageous Californians", "win"})

	league = append(league, leagues{"Allegoric Alaskans", 0, 0, 0, 0, 0})
	league = append(league, leagues{"Blithering Badgers", 0, 0, 0, 0, 0})
	league = append(league, leagues{"Courageous Californians", 0, 0, 0, 0, 0})
	league = append(league, leagues{"Devastating Donkeys", 0, 0, 0, 0, 0})

	for k, _ := range result {
		for x, _ := range league {
			if result[k].home == league[x].team {
				league[x].MP = league[x].MP + 1

				if result[k].outcome == WIN {
					league[x].W = league[x].W + 1
					league[x].P = league[x].P + 3

				} else if result[k].outcome == DRAW {
					league[x].D = league[x].D + 1
					league[x].P = league[x].P + 1
				} else {
					league[x].L = league[x].L + 1
				}
			}
		}
		for x, _ := range league {
			if result[k].away == league[x].team {
				league[x].MP = league[x].MP + 1

				if result[k].outcome == WIN {
					league[x].L = league[x].L + 1

				} else if result[k].outcome == DRAW {
					league[x].D = league[x].D + 1
					league[x].P = league[x].P + 1
				} else {
					league[x].W = league[x].W + 1
					league[x].P = league[x].P + 3
				}
			}
		}

	}

	sort.Slice(league, func(i, j int) bool {
		return league[i].P > league[j].P
	})

	fmt.Printf("Team                     | MP|  W|  D|  L|  P\n")
	for z, _ := range league {
		fmt.Printf("%-25s|%3d|%3d|%3d|%3d|%3d\n", league[z].team, league[z].MP, league[z].W, league[z].D, league[z].L, league[z].P)
	}
}
