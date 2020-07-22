package main

import (
	"fmt"
)

type games struct {
	ball1   int
	ball2   int
	fill    int
	score   int
	outcome string
}

const ends = 10
const maxscore = 10
const strike = "Strike"
const spare = "Spare"
const open = "Open"

var game [ends]games
var rollno int
var endno int

func main() {
	var game [ends]games
	roll(10, &game[endno])
	endno++
	roll(7, &game[endno])
	roll(2, &game[endno])
	endno++
	roll(8, &game[endno])
	roll(2, &game[endno])
	endno++
	roll(1, &game[endno])
	roll(6, &game[endno])
	endno++
	roll(9, &game[endno])
	roll(0, &game[endno])
	endno++
	roll(8, &game[endno])
		roll(1, &game[endno])
	endno++
	roll(0, &game[endno])
		roll(10, &game[endno])
	endno++
	roll(6, &game[endno])
		roll(2, &game[endno])
	endno++
	roll(10, &game[endno])
	endno++
	roll(10, &game[endno])
	roll(3, &game[endno])
	roll(10, &game[endno])
	finalscore := score(game)
	fmt.Println(finalscore)
}

func roll(pins int, game *games) {
	if game.outcome == "" {
		game.ball1 = pins
		game.score = pins
		if game.score == maxscore {
			game.outcome = strike
		} else {
			game.outcome = "Ball1"
		}
	} else {
		if endno == 9 && game.ball1 == 10 && game.ball2 != 0 {
			game.fill = pins
		} else {
			game.ball2 = pins
			game.score = pins + game.ball1
			if game.score == maxscore {
				game.outcome = spare
			} else {
				game.outcome = open
			}
		}
	}

//	fmt.Printf("In func roll endno = %v, outcome = %v, Ball1 = %v, Ball2 = %v, Fill = %v, pins = %v\n", endno, game.outcome, game.ball1, game.ball2, game.fill, pins)

}

func score(game [10]games) int {
	finalscore := 0

	for i := 0; i < ends; i++ {
		if game[i].outcome == spare {
			game[i].score = game[i].score + game[i+1].ball1
		}
		if game[i].outcome == strike {
			game[i].score = game[i].score + game[i+1].ball1
			if game[i+1].outcome == strike {
				game[i].score = game[i].score + game[i+2].ball1
			} else {
				game[i].score = game[i].score + game[i+1].ball2
			}
		}
		if i == 9 {
			game[i].score = game[i].score + game[i].fill
		}
	}

	fmt.Println(game)

	for i := 0; i < ends; i++ {
		finalscore += game[i].score
	}

	return finalscore
}

