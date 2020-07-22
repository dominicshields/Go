package main
// http://exercism.io/exercises/go/house/readme
import (
	"fmt"
)

func main() {
	//	const ARRMAX = 11
	this := "This is the "
	that := "that "
	the := " the "
	var objectfwd = []string{"house that Jack built.", "malt", "rat", "cat", "dog", "cow with the crumpled horn", "maiden all forlorn", "man all tattered and torn", "priest all shaven and shorn", "rooster that crowed in the morn", "farmer sowing his corn", "horse and the hound and the horn"}
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

	for x := 0; x < len(object); x++ {

		fmt.Printf("%v%v\n", this, objectfwd[x])

		for i := (len(object) - x); i < len(object); i++ {
			fmt.Printf("%v%v%v%v\n", that, action[i], the, object[i])
		}
		fmt.Printf("\n")
	}
}
