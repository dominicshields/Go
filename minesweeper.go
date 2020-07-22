package main

import (
	"fmt"
	"strconv"
)

const matrixrows = 6
const matrixcols = 7
const star = "*"

var matrix [matrixrows][matrixcols]string

func main() {
	starcount := 0
	setupmatrix()
	printmatrix()

	for i := 0; i < matrixrows; i++ {
		for j := 0; j < matrixcols; j++ {
			if matrix[i][j] == " " {
				starcount = 0
				for x := -1; x <= 1; x++ {
					for y := -1; y <= 1; y++ {
						if matrix[i+x][j+y] == star {
							starcount++
						}
					}
				}
				if starcount > 0 {
					matrix[i][j] = strconv.Itoa(starcount)
				}
			}
		}
	}

	printmatrix()

}

func printmatrix() {
	for i := 0; i < matrixrows; i++ {
		for j := 0; j < matrixcols; j++ {
			fmt.Printf("%v", matrix[i][j])
		}
		fmt.Printf("\n")
	}
}

func setupmatrix() {
	matrix[0][0] = "+"
	matrix[0][1] = "-"
	matrix[0][2] = "-"
	matrix[0][3] = "-"
	matrix[0][4] = "-"
	matrix[0][5] = "-"
	matrix[0][6] = "+"
	matrix[1][0] = "|"
	matrix[1][1] = " "
	matrix[1][2] = star
	matrix[1][3] = " "
	matrix[1][4] = star
	matrix[1][5] = " "
	matrix[1][6] = "|"
	matrix[2][0] = "|"
	matrix[2][1] = " "
	matrix[2][2] = " "
	matrix[2][3] = star
	matrix[2][4] = " "
	matrix[2][5] = " "
	matrix[2][6] = "|"
	matrix[3][0] = "|"
	matrix[3][1] = " "
	matrix[3][2] = " "
	matrix[3][3] = star
	matrix[3][4] = " "
	matrix[3][5] = " "
	matrix[3][6] = "|"
	matrix[4][0] = "|"
	matrix[4][1] = " "
	matrix[4][2] = " "
	matrix[4][3] = " "
	matrix[4][4] = " "
	matrix[4][5] = " "
	matrix[4][6] = "|"
	matrix[5][0] = "+"
	matrix[5][1] = "-"
	matrix[5][2] = "-"
	matrix[5][3] = "-"
	matrix[5][4] = "-"
	matrix[5][5] = "-"
	matrix[5][6] = "+"
}
