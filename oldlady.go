package main

import (
	"fmt"
)

func main() {

	intro := "I know an old lady who swallowed a "
	swallow := "She swallowed the "
	catch := " to catch the "
	spider := " that wriggled and jiggled and tickled inside her."
	var animal = []string{"fly", "spider", "bird", "cat", "dog", "goat", "cow", "horse"}
	var action = []string{"I don't know why she swallowed the fly. Perhaps she'll die.",
		"It wriggled and jiggled and tickled inside her.",
		"How absurd to swallow a bird!",
		"Imagine that, to swallow a cat!",
		"What a hog, to swallow a dog!",
		"Just opened her throat and swallowed a goat!",
		"I don't know how she swallowed a cow!",
		"She's dead, of course!"}

	for x, _ := range animal {
		fmt.Printf("%v%v\n%v\n", intro, animal[x], action[x])
		if x+1 == len(animal) {
			break
		}

		if x > 0 {
			fmt.Printf("%v%v%v%v", swallow, animal[x], catch, animal[x-1])
			if x == 2 {
				fmt.Printf("%v", spider)
			}
			fmt.Printf("\n")
			for i := x - 1; i > 0; i-- {
				fmt.Printf("%v%v%v%v", swallow, animal[i], catch, animal[i-1])
				if i == 2 {
					fmt.Printf("%v", spider)
				}
				fmt.Printf("\n")
			}
		}
		if x >= 1 {
			fmt.Printf("%v\n", action[0])
		}

		fmt.Printf("\n")
	}

}
