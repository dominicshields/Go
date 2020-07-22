package main

import (
	"fmt"
	"math"
)

type prime struct {
	num     float64
	isprime bool
}

func main() {
	const maxnumber = 10000
	var primes []prime
	var i float64
	cnt := 0

	for i = 2; i < maxnumber; i++ {
		primes = append(primes, prime{i, true})
	}

	for x := range primes {
		for t := 2.0; t <= math.Sqrt(primes[x].num); t++ {
			if math.Mod(primes[x].num, t) == 0 {
				primes[x].isprime = false
			}
		}
	}

	for x := range primes {
		if primes[x].isprime {
			fmt.Printf("%v\n", primes[x].num)
			cnt++
		}
	}
	fmt.Printf("Total number of prime numbers up to %v = %v\n", maxnumber, cnt)
}
