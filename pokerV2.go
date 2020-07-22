package main

import (
	"fmt"
	"math/rand"
	"sort"
	//	"time"
)

const cardmax = 13
const cardno = 5
const suitmax = 4
const players = 4

var suits map[int]string
var cards map[int]string
var handrank map[int]string
var encountered map[string]bool

type hands struct {
	cardval   [cardno]string
	cardidx   [cardno]int
	suit      [cardno]string
	suitidx   [cardno]int
	handvalue int
	highcard  int
}

func main() {
	suits = map[int]string{
		1: "Hearts",
		2: "Diamonds",
		3: "Clubs",
		4: "Spades",
	}
	cards = map[int]string{
		1:  "Ace",
		2:  "Two",
		3:  "Three",
		4:  "Four",
		5:  "Five",
		6:  "Six",
		7:  "Seven",
		8:  "Eight",
		9:  "Nine",
		10: "Ten",
		11: "Jack",
		12: "Queen",
		13: "King",
	}

	handrank = map[int]string{
		9: "Straight flush",
		8: "Four of a kind",
		7: "Full house",
		6: "Flush",
		5: "Straight",
		4: "Three of a kind",
		3: "Two pair",
		2: "One pair",
		1: "High card",
	}

	var hand [players]hands
	cardrand := 0
	suitrand := 0
	highesthandval := 0
	highestcardval := 0
	encountered := map[string]bool{}
	var highestcardidxarr []int
	var highesthandidxarr []int

	//	s1 := rand.NewSource(time.Now().UnixNano())
	s1 := rand.NewSource(147)
	r1 := rand.New(s1)

	for i := 0; i < players; i++ {
		for z := 0; z < cardno; z++ {
			cardrand = r1.Intn(cardmax) + 1
			suitrand = r1.Intn(suitmax) + 1

			key := cards[cardrand] + suits[suitrand]
			if checkDuplicates(key, encountered) == true {
				z--
			} else {
				hand[i].cardval[z] = cards[cardrand]
				hand[i].cardidx[z] = cardrand
				hand[i].suit[z] = suits[suitrand]
				hand[i].suitidx[z] = suitrand
			}

		}
		analyse(&hand[i])
	}

	for i := 0; i < players; i++ {
		fmt.Printf("\nPlayer %v  ", i+1)
		for z := 0; z < cardno; z++ {
			fmt.Printf("%v of %v    ", hand[i].cardval[z], hand[i].suit[z])
		}
	}
	fmt.Printf("\n")
	for i := 0; i < players; i++ {
		if hand[i].handvalue > highesthandval {
			highesthandval = hand[i].handvalue
		}
	}

	if highesthandval > 1 {
		for i := 0; i < players; i++ {
			if hand[i].handvalue == highesthandval {
				highesthandidxarr = append(highesthandidxarr, i)
			}
		}
	}

	if highesthandval == 1 {
		for i := 0; i < players; i++ {
			if hand[i].highcard > highestcardval {
				highestcardval = hand[i].highcard
			}
		}
		for i := 0; i < players; i++ {
			if hand[i].highcard == highestcardval {
				highestcardidxarr = append(highestcardidxarr, i)
			}
		}

		for i := 0; i < len(highestcardidxarr); i++ {
			fmt.Printf("Winning Hand = %v %v in hand %v\n", handrank[highesthandval], cards[highestcardval], hand[highestcardidxarr[i]])
		}

	} else {
		for i := 0; i < len(highesthandidxarr); i++ {
			fmt.Printf("Winning Hand = %v in hand %v\n", handrank[highesthandval], hand[highesthandidxarr[i]])
		}
	}
}

func analyse(hand *hands) {
	var cardcount [cardmax + 1]int // card has no zero but arrays do
	var suitcount [suitmax + 1]int
	var cardsarr []int
	var straight bool = true
	var assigned bool = false

	twoflag := 0
	threeflag := 0

	for i := 0; i < cardno; i++ {
		cardcount[hand.cardidx[i]]++
		suitcount[hand.suitidx[i]]++
		cardsarr = append(cardsarr, hand.cardidx[i])
	}

	sort.Ints(cardsarr)

	for i := 0; i < len(cardcount); i++ {
		if cardcount[i] == 4 {
			hand.handvalue = 8 // Four of a kind
			hand.highcard = i
			assigned = true
		}
		if cardcount[i] == 3 {
			threeflag = 1
			hand.handvalue = 4 // Three of a kind
			hand.highcard = i
			assigned = true
		}
		if cardcount[i] == 2 {
			twoflag++
			hand.highcard = i
			assigned = true
		}
	}

	if twoflag == 1 {
		hand.handvalue = 2 // Pair
	} else if twoflag == 2 {
		hand.handvalue = 3 // Two Pairs
	}

	if twoflag == 1 && threeflag == 1 {
		hand.handvalue = 7 // Full House
		hand.highcard = cardsarr[4]
	}

	for i := 0; i < len(cardsarr)-1; i++ {
		if cardsarr[i]+1 != cardsarr[i+1] {
			straight = false
		}
	}

	if straight {
		hand.handvalue = 5 // straight
		hand.highcard = cardsarr[4]
		assigned = true
	}

	for i := 0; i < len(suitcount); i++ {
		if suitcount[i] == 5 {
			if straight {
				hand.handvalue = 9 // Straight Flush
			} else {
				hand.handvalue = 6 // Flush
				hand.highcard = cardsarr[4]
				assigned = true
			}
		}
	}

	if !assigned {
		hand.handvalue = 1
		hand.highcard = cardsarr[4]
	}

}

func checkDuplicates(key string, encountered map[string]bool) bool {
	var retval bool
	if encountered[key] == true {
		retval = true
	} else {
		encountered[key] = true
		retval = false
	}
	return retval
}
