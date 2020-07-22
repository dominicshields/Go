package main

import (
	"fmt"
	"math"
)

func main() {
	const dot = "."
	const message = "WE ARE DISCOVERED FLEE AT ONCE"
	var rail1 []string
	var rail2 []string
	var rail3 []string
	var ciphertext []string
	var ciphertext1 []string
	var ciphertext2 []string
	var ciphertext3 []string
	cnt := 1
	flag := "UP"

	for x := range message {
		if string(message[x]) != " " {
			switch cnt {
			case 1:
				rail1 = append(rail1, string(message[x]))
				ciphertext1 = append(ciphertext1, string(message[x]))
				rail2 = append(rail2, dot)
				rail3 = append(rail3, dot)
				cnt = 2
				flag = "UP"
			case 2:
				rail1 = append(rail1, dot)
				rail2 = append(rail2, string(message[x]))
				ciphertext2 = append(ciphertext2, string(message[x]))
				rail3 = append(rail3, dot)
				if flag == "UP" {
					cnt = 3
				} else {
					cnt = 1
				}
			case 3:
				rail1 = append(rail1, dot)
				rail2 = append(rail2, dot)
				rail3 = append(rail3, string(message[x]))
				ciphertext3 = append(ciphertext3, string(message[x]))
				cnt = 2
				flag = "DOWN"
			}
		}
	}
	ciphertext = append(ciphertext1, ciphertext2...)
	ciphertext = append(ciphertext, ciphertext3...)
	fmt.Println(rail1)
	fmt.Println(rail2)
	fmt.Println(rail3)
	fmt.Println(ciphertext)

	rail1 = rail1[:0]
	rail2 = rail2[:0]
	rail3 = rail3[:0]

	for cnt = 1; cnt <= len(ciphertext); cnt++ {
		rail1 = append(rail1, dot)
		rail2 = append(rail2, dot)
		rail3 = append(rail3, dot)
	}

	s := 0
	for cnt = 0; cnt <= len(ciphertext); cnt++ {
		if math.Mod(float64(cnt), 4.0) == 0 {
			rail1[cnt] = ciphertext[s]
			s++
		}
	}

	for cnt = 1; cnt <= len(ciphertext); cnt++ {
		if math.Mod(float64(cnt), 2.0) == 0 {
			rail2[cnt-1] = ciphertext[s]
			s++
		}
	}

	for cnt = 1; cnt <= len(ciphertext); cnt++ {
		if math.Mod(float64(cnt), 4.0) == 0 {
			rail3[cnt-2] = ciphertext[s]
			s++
		}
	}

	fmt.Println(rail1)
	fmt.Println(rail2)
	fmt.Println(rail3)
}
