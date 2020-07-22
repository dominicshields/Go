package main

import (
	"fmt"
)

func main() {
	white := [2]int{7, 1}
	black := [2]int{5, 1}
	diagonals(white, black)
	rows(white, black)
	columns(white, black)

}

func diagonals(white [2]int, black [2]int) {
	white0 := white[0]
	white1 := white[1]
	for i := 0; i < 8; i++ {
		if white[0] == black[0] && white[1] == black[1] {
			fmt.Println("A. Queens can attack along diagonal")
		}
		white[0]++
		white[1]++
	}

	white[0] = white0
	white[1] = white1

	for i := 0; i < 8; i++ {
		if white[0] == black[0] && white[1] == black[1] {
			fmt.Println("B. Queens can attack along diagonal")
		}
		white[0]--
		white[1]--
	}

	white[0] = white0
	white[1] = white1
}

func rows(white [2]int, black [2]int) {
	white1 := white[1]
	for i := 0; i < 8; i++ {
		if white[0] == black[0] && white[1] == black[1] {
			fmt.Println("C. Queens can attack along row")
		}
		white[1]++
	}

	white[1] = white1

	for i := 0; i < 8; i++ {
		if white[0] == black[0] && white[1] == black[1] {
			fmt.Println("D. Queens can attack along row")
		}
		white[1]--
	}
}

func columns(white [2]int, black [2]int) {
	white0 := white[0]
	for i := 0; i < 8; i++ {
		if white[0] == black[0] && white[1] == black[1] {
			fmt.Println("E. Queens can attack along column")
		}
		white[0]++
	}

	white[0] = white0

	for i := 0; i < 8; i++ {
		if white[0] == black[0] && white[1] == black[1] {
			fmt.Println("F. Queens can attack along column")
		}
		white[0]--
	}
}
