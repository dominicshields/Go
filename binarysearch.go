package main

import (
	"fmt"
//	"math"
	"sort"
)

var loopcnt int
var lower int
var upper int

func main() {
	const searchpattern = "Peter A Wright/TITCHFIELD/ONS"
	employee := []string{}
	employee = append(employee, "TONY LIU/NEWPORT/ONS")
	employee = append(employee, "Rachel Varney/TITCHFIELD/ONS")
	employee = append(employee, "Madara Watkins/NEWPORT/ONS")
	employee = append(employee, "Ian Cullen/NEWPORT/ONS")
	employee = append(employee, "Tanya Vaughan/NEWPORT/ONS")
	employee = append(employee, "Marianna Vasili/TITCHFIELD/ONS")
	employee = append(employee, "Kate Vaughan Jones/NEWPORT/ONS")
	employee = append(employee, "Cath Fawcett/NEWPORT/ONS")
	employee = append(employee, "David J Williams/NEWPORT/ONS")
	employee = append(employee, "Ben Whitestone/NEWPORT/ONS")
	employee = append(employee, "Penny Witham/NEWPORT/ONS")
	employee = append(employee, "George Zygmund/NEWPORT/ONS")
	employee = append(employee, "Peter A Wright/TITCHFIELD/ONS")
	employee = append(employee, "Lewis Hodge/NEWPORT/ONS")
	employee = append(employee, "Sharne Bailey/NEWPORT/ONS")
	employee = append(employee, "Nick Vaughan/NEWPORT/ONS")
	employee = append(employee, "Mike Wendes/TITCHFIELD/ONS")
	employee = append(employee, "Nikola Bowers/TITCHFIELD/ONS")
	employee = append(employee, "Donna Wall/NEWPORT/ONS")
	employee = append(employee, "Sreeram Vedagiri/Titchfield/ONS")
	employee = append(employee, "Stacey Vaughan/NEWPORT/ONS")
	employee = append(employee, "Stephen D King/TITCHFIELD/ONS")
	employee = append(employee, "Julia Abbott/NEWPORT/ONS")
	employee = append(employee, "Ruth Williams/NEWPORT/ONS")
	employee = append(employee, "Aaron Aardvark/NEWPORT/ONS")

	sort.Strings(employee)
	split := len(employee) / 2
	upper = len(employee)
	lower = 0
	chop(employee, split, searchpattern)

}
func chop(employee []string, split int, searchpattern string) bool {
	loopcnt++
	fmt.Printf("Split = %v, Upper = %v, Lower = %v\n", split, upper, lower)
	if loopcnt > len(employee)/2 {
		fmt.Printf("Search Pattern %v is not found in %v iterations. Split = %v\n", searchpattern, loopcnt, split)
	} else if employee[split] > searchpattern {
		if split < upper {
			upper = split
		}
		//		split = int(math.Floor(float64((upper - lower) / 2)))
		split -= (upper - lower) / 2
		split-- //				Because array starts at zero and we would never get there otherwise
		chop(employee, split, searchpattern)
	} else if employee[split] < searchpattern {
		if split > lower {
			lower = split
		}
		split += (upper - lower) / 2
		chop(employee, split, searchpattern)
	} else {
		fmt.Printf("Found = %v at array position %v in %v iterations\n", employee[split], split, loopcnt)
	}
	return true
}
