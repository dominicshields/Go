package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func main() {
	const lower = 10
	const upper = 99
	var products []int
	var palindromes []int
	var lowerfactors []int
	var upperfactors []int
	var producthold string
	var productreverse string

	for i := lower; i <= upper; i++ {
		for j := lower; j <= upper; j++ {
			products = append(products, i*j)
		}
	}
	products = removeDuplicates(products)

	for x := range products {
		producthold = strconv.Itoa(products[x])
		data := []rune(producthold)
		result := []rune{}
		for z := len(data) - 1; z >= 0; z-- {
			result = append(result, data[z])
		}
		productreverse = string(result)
		if producthold == productreverse {
			pal, _ := strconv.Atoi(producthold)
			palindromes = append(palindromes, pal)
		}
	}

	sort.Slice(palindromes, func(i, j int) bool {
		return palindromes[i] < palindromes[j]
	})

	for y := 1; y <= palindromes[0]; y++ {
		if math.Mod(float64(palindromes[0]), float64(y)) == 0 {
			lowerfactors = append(lowerfactors, y)
		}
	}

	for y := 1; y <= palindromes[len(palindromes)-1]; y++ {
		if math.Mod(float64(palindromes[len(palindromes)-1]), float64(y)) == 0 {
			upperfactors = append(upperfactors, y)
		}
	}

	fmt.Printf("Lowest = %v, Highest = %v\n", palindromes[0], palindromes[len(palindromes)-1])
	fmt.Printf("Factors of %v = %v\n", palindromes[0], lowerfactors)
	fmt.Printf("Factors of %v = %v\n", palindromes[len(palindromes)-1], upperfactors)
}

func removeDuplicates(a []int) []int {
	result := []int{}
	seen := map[int]int{}
	for _, val := range a {
		if _, ok := seen[val]; !ok {
			result = append(result, val)
			seen[val] = val
		}
	}
	return result
}
