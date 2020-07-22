package main

import (
	"fmt"
)

type saddles struct {
	value int
	row   int
	col   int
}

/*  0  1  2
  |---------
0 | 9  8  7
1 | 5  3  2     <--- saddle point at (1,0) It's called a "saddle point" because it is greater than or equal to every element in its row and less than or equal to every element in its column.
2 | 6  6  7  */

func main() {
	const matrixrows = 3
	const matrixcols = 3
	saddlerows := []saddles{}
	saddlecols := []saddles{}
	highestinrow := 0
	lowestincol := 99999
	rowstore := 0
	colstore := 0
	highflag := 0
	lowflag := 0
	var matrix [matrixrows][matrixcols]int
	matrix[0][0] = 9
	matrix[0][1] = 8
	matrix[0][2] = 7
	matrix[1][0] = 5
	matrix[1][1] = 5
	matrix[1][2] = 2
	matrix[2][0] = 6
	matrix[2][1] = 6
	matrix[2][2] = 7

	fmt.Printf("Matrix = %v\n", matrix)

	for i := 0; i < matrixrows; i++ {
		highestinrow = 0
		for x := 0; x < matrixcols; x++ {
			highflag = 0
			if matrix[i][x] >= highestinrow {
				highestinrow = matrix[i][x]
				rowstore = i
				colstore = x
				highflag = 1
			}
			if highflag == 1 {
				saddlerows = append(saddlerows, saddles{highestinrow, rowstore, colstore})
			}
		}
	}

	//	fmt.Printf("saddlerows = %v\n", saddlerows)

	for i := 0; i < matrixcols; i++ {
		lowestincol = 99999
		for x := 0; x < matrixrows; x++ {
			lowflag = 0
			if matrix[x][i] <= lowestincol {
				lowestincol = matrix[x][i]
				colstore = x
				rowstore = i
				lowflag = 1
			}
			if lowflag == 0 {
				saddlecols = append(saddlecols, saddles{lowestincol, rowstore, colstore})
			}
		}
	}
	//	fmt.Printf("saddlecols = %v\n", saddlecols)

	for i := 0; i < len(saddlerows); i++ {
		for x := 0; x < len(saddlecols); x++ {
			if saddlerows[i].value == saddlecols[x].value && saddlerows[i].row == saddlecols[x].col && saddlerows[i].col == saddlecols[x].row {
				fmt.Printf("Saddle Point = Value %v, Row %v, Col %v\n", saddlerows[i].value, saddlerows[i].row, saddlerows[i].col)
			}
		}
	}

}
