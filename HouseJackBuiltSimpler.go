package main

// http://exercism.io/exercises/go/house/readme
import (
	"fmt"
)

func main() {
	this := "This is the "
	that := "that "
	the := " the "
	var object = []string{
		"horse and the hound and the horn",
		"farmer sowing his corn",
		"rooster that crowed in the morn",
		"priest all shaven and shorn",
		"man all tattered and torn",
		"maiden all forlorn",
		"cow with the crumpled horn",
		"dog",
		"cat",
		"rat",
		"malt",
		"house that Jack built."}
	var action = []string{"", "belonged to", "kept", "woke", "married", "kissed", "milked", "tossed", "worried", "killed", "ate", "lay in"}

	for x := len(object) -1; x > -1; x-- {

		fmt.Printf("%v%v\n", this, object[x])

		for i := x +1 ; i < len(object); i++ {
			fmt.Printf("%v%v%v%v\n", that, action[i], the, object[i])
		}
		fmt.Printf("\n")
	}
}
