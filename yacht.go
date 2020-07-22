package main

/*
Category        Score                   Example
Ones            1 x number of ones      1 1 1 4 5 scores 3
Twos            2 x number of twos      2 2 3 4 5 scores 4
Threes          3 x number of threes    3 3 3 3 3 scores 15
Fours           4 x number of fours     1 2 3 3 5 scores 0
Fives           5 x number of fives     5 1 5 2 5 scores 15
Sixes           6 x number of sixes     2 3 4 5 6 scores 6
Full House      Total of the dice
Three of one number and two of another  3 3 3 5 5 scores 19
Four of a Kind  Total of the four dice  4 4 4 4 6 scores 16
Little Straight 30 points               1 2 3 4 5 scores 30
Big Straight    30 points               2 3 4 5 6 scores 30
Choice          Sum of the dice         2 3 3 4 6 scores 18
Yacht           50 points               4 4 4 4 4 scores 50
*/

import (
	"fmt"
	"math/rand"
	"time"
)

const dicemax = 6
const diceno = 5
const catno = 12
const straightscore = 30
const yachtscore = 50
const maxreps = 1000000

type scorehist struct {
	die [diceno]int
	cat string
	sco int
}

func main() {
	var category = []string{"Ones", "Twos", "Threes", "Fours", "Fives", "Sixes", "Full House", "Four of a Kind", "Little Straight", "Big Straight", "Choice", "Yacht"}
	var dice [diceno]int
	var scorehists []scorehist
	var topscorehists []scorehist
	catrand := 0
	score := 0
	highscore := 0
	lowscore := 999
	var totalscore int64

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for w := 0; w < maxreps; w++ {
		overallscore := 0

		for catrand = 0; catrand < len(category); catrand++ {
			for i := 0; i < diceno; i++ {
				dice[i] = r1.Intn(dicemax) + 1
			}
			//	catrand := r1.Intn(catno)    // In single use mode this needs to be moved up and enabled obv.

			switch category[catrand] {
			case "Ones":
				score = scorenum(dice, 1)
			case "Twos":
				score = scorenum(dice, 2)
			case "Threes":
				score = scorenum(dice, 3)
			case "Fours":
				score = scorenum(dice, 4)
			case "Fives":
				score = scorenum(dice, 5)
			case "Sixes":
				score = scorenum(dice, 6)
			case "Full House":
				score = fullhouse(dice)
			case "Four of a Kind":
				score = fourofakind(dice)
			case "Little Straight":
				score = straight(dice, 1)
			case "Big Straight":
				score = straight(dice, 2)
			case "Choice":
				score = choice(dice)
			case "Yacht":
				score = yacht(dice)
			}
			//	fmt.Printf("%v %v\tScore = %v\n", dice, category[catrand], score)
			overallscore += score
			scorehists = append(scorehists, scorehist{dice, category[catrand], score})
		}
		//fmt.Printf("Overall 12 Game \tScore = %v\n", overallscore)
		if overallscore > highscore {
			highscore = overallscore
			topscorehists = scorehists
		}
		if overallscore < lowscore {
			lowscore = overallscore
		}
		totalscore += int64(overallscore)
					scorehists = nil
	}
	fmt.Printf("High Score = %v, Low Score = %v, Average Score = %v\n", highscore, lowscore, totalscore/maxreps)
	for _,v := range topscorehists {
			fmt.Printf("%v\n",v)
	}

}

/////////////////////////
func scorenum(dice [diceno]int, number int) int {
	score := 0
	times := 0
	for i := 0; i < len(dice); i++ {
		if dice[i] == number {
			times++
		}
	}
	score = times * number
	return score
}

/////////////////////////
func fullhouse(dice [diceno]int) int {
	var dicescore [dicemax + 1]int // dice has no zero but arrays do
	score := 0
	threeflag := 0
	twoflag := 0
	for i := 0; i < len(dice); i++ {
		dicescore[dice[i]]++
		score += dice[i]
	}

	for i := 0; i < len(dicescore); i++ {
		if dicescore[i] == 3 {
			threeflag = 1
		}
		if dicescore[i] == 2 {
			twoflag = 1
		}
	}

	if threeflag == 1 && twoflag == 1 {
		return score
	} else {
		return 0
	}
}

/////////////////////////
func fourofakind(dice [diceno]int) int {
	var dicescore [dicemax + 1]int // dice has no zero but arrays do
	score := 0
	fourflag := 0
	for i := 0; i < len(dice); i++ {
		dicescore[dice[i]]++
	}
	for i := 0; i < len(dicescore); i++ {
		if dicescore[i] == 4 {
			fourflag = 1
			score = 4 * i
		}
	}

	if fourflag == 1 {
		return score
	} else {
		return 0
	}
}

/////////////////////////
func straight(dice [diceno]int, lower int) int {
	var dicescore [dicemax + 1]int // dice has no zero but arrays do
	flag := 0
	for i := 0; i < len(dice); i++ {
		dicescore[dice[i]]++
	}
	for i := 0; i < len(dicescore); i++ {
		if dicescore[i] == 1 {
			flag++
			if lower == 1 && i == 6 {
				flag--
			}
			if lower == 2 && i == 1 {
				flag--
			}
		}
	}

	if flag == 5 {
		return straightscore
	} else {
		return 0
	}
}

/////////////////////////
func choice(dice [diceno]int) int {
	score := 0
	for i := 0; i < len(dice); i++ {
		score += dice[i]
	}
	return score
}

/////////////////////////
func yacht(dice [diceno]int) int {
	score := 0
	times := 0
	firstdice := dice[0]
	for i := 0; i < len(dice); i++ {
		if dice[i] == firstdice {
			times++
		}
	}
	if times == 5 {
		score = yachtscore
	} else {
		score = 0
	}
	return score
}
