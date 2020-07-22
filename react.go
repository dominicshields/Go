package main

import (
	"fmt"
)

type buffers struct {
	value string
	age   int
	stale bool
	react int
}

const ARRMAX = 7 // Arbitrary buffer size

var ageval int

func main() {
	const newinput = "Z"
	var buffer = [ARRMAX]buffers{}
	var alphabet []string
	ageval = 1

	for i := 1; i <= ARRMAX; i++ {
		alphabet = append(alphabet, string('a'-1+i))
	}

	for x := range buffer {
		buffer[x].value = alphabet[x]
		buffer[x].age = ageval
		buffer[x].stale = false
		ageval++
	}

	alphabet = alphabet[:0]
	for i := 1; i <= ARRMAX; i++ {
		alphabet = append(alphabet, string('j'-1+i))
	}
	fmt.Println(buffer)
	for x := range buffer {
		if !findspace(buffer) {
			buffer = removeoldest(buffer)
		}
		buffer = writebuffer(buffer, alphabet[x])
		fmt.Println(buffer)
		buffer = react(buffer)
	}
}

func react(buffer [ARRMAX]buffers) [ARRMAX]buffers {
	//	f := 0
	for i := 0; i < len(buffer); i++ {
		if buffer[i].stale == true {
			fmt.Printf("Field updated: %v\n", buffer[i].value)
			buffer = reactoperation(buffer, i)
			buffer[i].stale = false
		}
	}
	return buffer
}

func reactoperation(buffer [ARRMAX]buffers, i int) [ARRMAX]buffers {
	buffer[i].react = buffer[i].age * ARRMAX
	return buffer
}

func writebuffer(buffer [ARRMAX]buffers, new string) [ARRMAX]buffers {
	for i := 0; i < len(buffer); i++ {
		if buffer[i].age == 0 {
			buffer[i].age = ageval
			buffer[i].value = new
			buffer[i].stale = true
			break
		}
	}
	ageval++
	return buffer
}

func findspace(buffer [ARRMAX]buffers) bool {
	findspace := false
	for i := 0; i < len(buffer); i++ {
		if buffer[i].age == 0 {
			findspace = true
			break
		}
	}
	return findspace
}

func removeoldest(buffer [ARRMAX]buffers) [ARRMAX]buffers {
	oldestidx := 9999999
	arridx := 0
	for i := 0; i < len(buffer); i++ {
		if buffer[i].age < oldestidx {
			oldestidx = buffer[i].age
			arridx = i
		}
	}
	buffer[arridx].age = 0
	return buffer
}
