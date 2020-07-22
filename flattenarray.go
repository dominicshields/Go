package main

import "fmt"

func main() {
	// Create an empty slice of slices.
	animals := [][]string{}
	flatslice := []string{}

	// Create three string slices.
	row1 := []string{"fish", "shark", "eel", "null"}
	row2 := []string{"bird"}
	row3 := []string{"lizard", "salamander"}

	// Append string slices to outer slice.
	animals = append(animals, row1)
	animals = append(animals, row2)
	animals = append(animals, row3)

	for i := range animals {
		for x := range animals[i] {
			if animals[i][x] != "null" {
				flatslice = append(flatslice, animals[i][x])
			}
		}
	}
	fmt.Printf("%v\n", flatslice)
}
