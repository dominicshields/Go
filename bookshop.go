package main

import (
	"fmt"
)

/*
One copy of any of the five books costs $8.
If, however, you buy two different books, you get a 5% discount on those two books.
If you buy 3 different books, you get a 10% discount.
If you buy 4 different books, you get a 20% discount.
If you buy all 5, you get a 25% discount.
Note: that if you buy four books, of which 3 are different titles, you get a 10% discount on the 3 that form part of a set, but the fourth book still costs $8.
*/
const singlebook = 8.0
const twobookdisc = 0.95
const threebookdisc = 0.9
const fourbookdisc = 0.8
const fivebookdisc = 0.75
const titles = 5

type bookinvs struct {
	bookno    int
	bookcount int
}

func main() {
	var books = []int{1, 1, 2, 2, 2, 2, 3, 3, 4, 4, 5}
	fmt.Printf("Total cost = £%.2f\n", Cost(books))
}

func Cost(books []int) float64 {
	var bookinv [titles]bookinvs
	booknobought := 0
	booktotbought := 0
	totalcost := 0.0
	bookmin := 99

	for i := 0; i < titles; i++ {
		bookinv[i].bookno = i + 1
	}

	for i := 0; i < len(books); i++ {
		for x := 0; x < titles; x++ {
			if books[i] == bookinv[x].bookno {
				bookinv[x].bookcount++
				break
			}
		}
	}

	fmt.Printf("books = %v\n", books)
	fmt.Printf("bookinv = %v\n", bookinv)

	for i := 0; i < len(bookinv); i++ {
		if bookinv[i].bookcount > 0 {
			booknobought++
			booktotbought += bookinv[i].bookcount
		}
	}

	fmt.Printf("Different Books bought %v\n", booknobought)
	fmt.Printf("Total Books bought %v\n", booktotbought)

	for i := 0; i < len(bookinv); i++ {
		if bookinv[i].bookcount > 0 {
			if bookinv[i].bookcount < bookmin {
				bookmin = bookinv[i].bookcount
			}

		}
	}

	switch booknobought {
	case 0:
	case 1:
		totalcost = singlebook * float64(booktotbought)
	case 2:
		bookmin *= 2
		totalcost = float64(singlebook*float64(bookmin)) * twobookdisc
		totalcost += float64(singlebook * (float64(booktotbought) - float64(bookmin)))
	case 3:
		bookmin *= 3
		totalcost = float64(singlebook*float64(bookmin)) * threebookdisc
		totalcost += float64(singlebook * (float64(booktotbought) - float64(bookmin)))
	case 4:
		bookmin *= 4
		totalcost = float64(singlebook*float64(bookmin)) * fourbookdisc
		totalcost += float64(singlebook * (float64(booktotbought) - float64(bookmin)))
	case 5:
		bookmin *= 5
		totalcost = float64(singlebook*float64(bookmin)) * fivebookdisc
		totalcost += float64(singlebook * (float64(booktotbought) - float64(bookmin)))
	}
	return totalcost
}
