package main

import (
	"fmt"
	"strings"
	"time"
)

var months map[string]string
var weeks map[string]int
var days map[string]int

func main() {
	const datemeetup = "The last friday of February 2018"
	//		const datemeetup = "The wednesteenth of May 2018"
	weekint := 0
	startday := 0
	var day string
	months := map[string]int{
		"January":   1,
		"February":  2,
		"March":     3,
		"April":     4,
		"May":       5,
		"June":      6,
		"July":      7,
		"August":    8,
		"September": 9,
		"October":   10,
		"November":  11,
		"December":  12,
	}

	weeks := map[string]int{
		"first":        1,
		"second":       2,
		"third":        3,
		"fourth":       4,
		"fifth":        5,
		"last":         5,
		"monteenth":    3,
		"tuesteenth":   3,
		"wednesteenth": 3,
		"thursteenth":  3,
		"friteenth":    3,
		"saturteenth":  3,
		"sunteenth":    3,
	}

	days := map[string]int{
		"sun": 0,
		"mon": 1,
		"tue": 2,
		"wed": 3,
		"thu": 4,
		"fri": 5,
		"sat": 6,
	}

	yearidx := strings.LastIndex(datemeetup, " ") // 	Extract the year
	year := datemeetup[yearidx+1:]

	substr1 := datemeetup[4:yearidx] // 			Get hold of the sentence without the intro and year

	monthidx := strings.LastIndex(substr1, " ") // 		Extract the month
	month := substr1[monthidx+1:]
	monthint := months[month] // 				Convert month to integer

	weekidx := strings.Index(substr1, " ") //		Extract the week
	week := substr1[:weekidx]
	weekint = weeks[week] //				Convert week (or proxy) to integer, adjust for February
	if monthint == 2 {    // February cannot have five
		if weekint == 5 {
			weekint = 4
		}
	}

	if weekint != 3 && week != "third" { //			Extract the day making exception for the silly, made up days that are not enhancing this exercise at all
		substr2 := substr1[weekidx+1:]
		dayidx := strings.Index(substr2, " ")
		day = substr2[:dayidx]
	} else {
		day = week
	}
	weekday := days[strings.ToLower(day[:3])] //		Return day of week as a number

	switch weekint { //					Relate weeknumber to idealised calendar start day
	case 1:
		startday = 1
	case 2:
		startday = 8
	case 3:
		startday = 15
	case 4:
		startday = 22
	case 5:
		startday = 29
	}

	strStart := fmt.Sprintf("%02d-%02d-%04s", startday, months[month], year) // Initialise date to the appropriate year and month and to the day returned from the switch/case
	start, _ := time.Parse("02-01-2006", strStart)

	for d := 0; d < 7; d++ { //						Loop around incrementing days until the day of the week matches the input
		date := start.AddDate(0, 0, d)
		if date.Weekday() == time.Weekday(weekday) {
			if date.Month() != time.Month(monthint) {
				date = start.AddDate(0, 0, d-7)
			}
			fmt.Println(date.Format("2006-01-02"))
		}
	}
}
