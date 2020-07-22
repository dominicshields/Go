package main

import (
	"fmt"
	"math"
)

type allergies struct {
	allergy string
	value   int
}

func main() {
	allergy := []allergies{
		{"Null", 1024},
		{"Null", 512},
		{"Null", 256},
		{"cats", 128},
		{"pollen", 64},
		{"chocolate", 32},
		{"tomatoes", 16},
		{"strawberries", 8},
		{"shellfish", 4},
		{"peanuts", 2},
		{"eggs", 1},
	}
	score := 50
	rem := float64(score)
	allergyidx := 0

	for rem != 0 {
		if rem > 0 {
			allergyidx = findallergy(int(rem), allergy)
		}
		rem = math.Mod(float64(score), float64(allergy[allergyidx].value))
	}
}

func findallergy(score int, allergy []allergies) int {
	x := 0
	for x = 0; x < len(allergy); x++ {
		if score >= allergy[x].value {
			break
		}
	}
	if allergy[x].allergy != "Null" {
		fmt.Printf("Allergic to %v\n", allergy[x].allergy)
	}
	return x
}