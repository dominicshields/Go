package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	var robotnames []string
	var randomno string
	var flag int
	rand.Seed(time.Now().UTC().UnixNano())
	robotnames = append(robotnames, "AD444")
	for y := 0; y < 20; y++ {
		randomno = randomString(2) + strconv.Itoa(randInt(1, 999))
		flag = 0
		for x, _ := range robotnames {
			if robotnames[x] == randomno {
				fmt.Printf("Error %v already exists\n", randomno)
				flag = 1
				y--
			}
		}
		if flag == 0 {
			robotnames = append(robotnames, randomno)
		}
	}
	fmt.Printf("%v\n", robotnames)
}

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

