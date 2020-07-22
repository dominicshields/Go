package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("Hour Minute Second: ", now.Format("03:04:05 PM"))
	then := now.Add(-10 * time.Minute)
	fmt.Println("10 minutes ago:", then.Format("03:04:05 PM"))
	future := now.Add(10 * time.Minute)
	fmt.Println("10 minutes ahead:", future.Format("03:04:05 PM"))
}