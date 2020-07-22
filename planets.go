package main

import (
	"fmt"
)

type planet struct {
	name         string
	yearfraction float64
}

func main() {
	const earthsecs = 31557600
	const agesecs = 1000000000

	planetinst := []planet{}

	planetinst = append(planetinst, planet{"Mercury", 0.2408467})
	planetinst = append(planetinst, planet{"Venus", 0.61519726})
	planetinst = append(planetinst, planet{"Earth", 1.0000000})
	planetinst = append(planetinst, planet{"Mars", 1.8808158})
	planetinst = append(planetinst, planet{"Jupiter", 11.862615})
	planetinst = append(planetinst, planet{"Saturn", 29.447498})
	planetinst = append(planetinst, planet{"Uranus", 84.016846})
	planetinst = append(planetinst, planet{"Neptune", 164.79132})

	for i, _ := range planetinst {
		//		fmt.Printf("Struct contains %s   %v\n", planetinst[i].name, planetinst[i].yearfraction)
		fmt.Printf("On %s a person is %d seconds old at age %f\n", planetinst[i].name, agesecs, (agesecs/earthsecs)/planetinst[i].yearfraction)
	}
}