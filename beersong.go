package main

import (
	"fmt"
)

func main() {
	const beermax = 99
	const bottle = " bottle "
	const bottles = " bottles "
	const ofbeer = "of beer"
	const onthewall = " on the wall"
	const takeone = "Take one"
	const down = " down and pass it around, "
	const nomore = "no more"

	for i := beermax; i >= 0; i-- {
		switch {
		case i > 2:
			fmt.Printf("%v%v%v%v, %v%v%v.\n", i, bottles, ofbeer, onthewall, i, bottles, ofbeer)
			fmt.Printf("%v%v%v%v%v%v.\n\n", takeone,down, i-1, bottles, ofbeer, onthewall)
		case i == 2:
			fmt.Printf("%v%v%v%v, %v%v%v.\n", i, bottles, ofbeer, onthewall, i, bottles, ofbeer)
			fmt.Printf("%v%v%v%v%v%v.\n\n", takeone,down, i-1, bottle, ofbeer, onthewall)
		case i == 1:
			fmt.Printf("%v%v%v%v, %v%v%v.\n", i, bottle, ofbeer, onthewall, i, bottle, ofbeer)
			fmt.Printf("Take it%v%v%v%v%v.\n\n",down, nomore, bottles, ofbeer, onthewall)
		case i == 0:
			fmt.Printf("%v%v%v%v, %v%v%v.\nGo to the store and buy some more, %v%v%v%v.\n", nomore, bottles, ofbeer, onthewall, nomore, bottles, ofbeer, beermax, bottles, ofbeer, onthewall)
		}

	}
}
