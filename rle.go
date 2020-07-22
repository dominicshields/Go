package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//	const data = "WWWWWWWWWWWWBWWWWWWWWWWWWBBBWWWWWWWWWWWWWWWWWWWWWWWWB" // "12WB12W3B24WB"
	const data = "AABCCCDEEEE" //->  "2AB3CD4E"  ->  "AABCCCDEEEE"

	checker := RLEncode(data)
	fmt.Println(checker)

	unchecker := RLdecode(checker)
	fmt.Println(unchecker)
}
func RLEncode(data string) string {
	var checker string
	var store string
	count := 0
	for _, v := range data {
		if string(v) == store {
			count++
		} else {
			if count > 0 {
				if count > 1 {
					checker += strconv.Itoa(count) + store
				} else {
					checker += store
				}
			}
			store = string(v)
			count = 1
		}
	}

	if count > 1 {
		checker += strconv.Itoa(count) + store
	} else {
		checker += store
	}
	return checker
}

func RLdecode(data string) string {
	var original string
	var charnum string
	var let string
	charnumlen := 0

	for _, v := range data {
		if _, err := strconv.Atoi(string(v)); err == nil {
			charnum += string(v)
		} else {
			let = string(v)
			if charnum != "" {
				charnumlen, _ = strconv.Atoi(charnum)
			} else {
				charnumlen = 1
			}
			original += strings.Repeat(let, charnumlen)
			charnum = ""
		}
	}
	return original
}
