package main

import (
	"fmt"
)

func main() {
	const STRINGVAL = "1234567890AB"
	const NUMCHARS = 6
	var substrings []string
	var unsafe string
	var ret bool

	substrings = All(NUMCHARS, STRINGVAL)
	fmt.Println(substrings)

	unsafe, ret = UnsafeFirst(NUMCHARS, STRINGVAL)
	if ret == true {
		fmt.Println(unsafe)
	} else {
		fmt.Println("String is too short")
	}
}

func All(n int, s string) []string { //	All returns a list of all substrings of s with length n.
	var allstrings []string
	var found string

	for i, c := range s {
		if (i+1)%n == 0 { // fiddle as i starts at zero
			found = found + string(c)
			allstrings = append(allstrings, found)
			found = ""
		} else {
			found = found + string(c)
		}
	}

	return allstrings
}

func UnsafeFirst(n int, s string) (unsafestring string, ok bool) { //UnsafeFirst returns the first substring of s with length n.
	if len(s) >= n {
		unsafestring = s[0:n]
		ok = true
	} else {
		ok = false
	}
	return unsafestring, ok

}
