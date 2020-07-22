package main

import (
	"fmt"
)

func main() {
	const daysecs = 86400
	const yeardays = 365.25
	const gigasecs = 1000000000

	fmt.Printf("A person is %d seconds old at age %f", gigasecs, gigasecs/(daysecs*yeardays))
}
