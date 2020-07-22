package main

import (
	"fmt"
	"strconv"
)

func main() {
	const base10num = 11371
	const divisor = 2
	const bitlen = 7
	var binarystring string
	var lsb string
	var vlq string
	rem := 0

	for numerator := base10num; numerator > 0; {
		rem = numerator % divisor
		binarystring = strconv.Itoa(rem) + binarystring
		numerator /= divisor
	}
	fmt.Printf("Decimal number = %v\n", base10num)
	fmt.Printf("binarystring = %v\n", binarystring)
	binarylen := len(binarystring)

	if binarylen < bitlen {
		lsb = binarystring
		for i := len(lsb); i < bitlen; i++ {
			lsb = "0" + lsb
		}
		lsb = "0" + lsb
		vlq = lsb
	} else {
		lsb = binarystring[len(binarystring)-bitlen:]
		for i := len(lsb); i < bitlen; i++ {
			lsb = "0" + lsb
		}
		lsb = "0" + lsb
		msb := binarystring[0 : len(binarystring)-bitlen]
		for i := len(msb); i < bitlen; i++ {
			msb = "0" + msb
		}
		msb = "1" + msb
		vlq = msb + lsb
	}
	fmt.Printf("vlq = %v\n", vlq)
}
