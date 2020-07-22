package main

import (
	"fmt"
	"unicode"
)

const maxlen = 10

func main() {
	const num1 = "+1 (613)-995-0253"
	const num2 = "613-995-0253"
	const num3 = "1 613 995 0253"
	const num4 = "613.995.0253"

	fmt.Println("Number: " + Number(num3))
	fmt.Println("Formatted: " + format(num4))
	fmt.Println("Area Code: " + AreaCode(num1))

}
func Number(num string) string {
	var clean string
	for _, y := range num {
		if unicode.IsNumber(y) {
			clean += string(y)
		}
	}
	if len(clean) > maxlen {
		clean = clean[len(clean)-maxlen:]
	}
	return clean
}
func format(num string) string {
	clean := Number(num)
	formatted := "(" + clean[0:3] + ") " + clean[3:6] + "-" + clean[6:10]
	return formatted
}
func AreaCode(num string) string {
	clean := Number(num)
	areacode := clean[0:3]
	return areacode
}
