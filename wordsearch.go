package main

import (
	"fmt"
	"sort"
	"strings"
)

const matrixrows = 10
const matrixcols = 10

var matrix [matrixrows][matrixcols]string
var allstrings []string

func main() {
	var wordsearch = []string{"jefblpepre", "camdcimgtc", "oivokprjsm", "pbwasqroua", "rixilelhrs", "wolcqlirpc", "screeaumgr", "alxhpburyi", "jalaycalmp", "clojurermt"}
	var words = []string{"clojure", "elixir", "ecmascript", "rust", "java", "lua", "lisp", "ruby"}
	//var words = []string{"clojure"}
	var letters string
	var reverse string
	setup(wordsearch)

	for y := 0; y < matrixrows; y++ { //			all words by row
		for z := 0; z < matrixcols; z++ { // jefblpepre erpeplbfej camdcimgtc ctgmicdmac oivokprjsm msjrpkovio pbwasqroua auorqsawbp rixilelhrs srhlelixir wolcqlirpc cprilqclow screeaumgr rgmuaeercs alxhpburyi iyrubphxla jalaycalmp pmlacyalaj clojurermt tmrerujolc
			letters += matrix[y][z]
		}
		allstrings = append(allstrings, letters)
		reverse = reversestring(letters)
		allstrings = append(allstrings, reverse)
		letters = ""
	}

	for y := 0; y < matrixcols; y++ { //			all words by column
		for z := 0; z < matrixrows; z++ { // jcoprwsajc cjaswrpocj eaibioclal lalcoibiae fmvwxlrxlo olxrlxwvmf bdoaicehaj jaheciaodb lckslqepyu uypeqlskcl pipqelabcr rcbaleqpip emrrliuuae eauuilrrme pgjohrmrlr rlrmrhojgp rtsurpgymm mmygprustr ecmascript tpircsamce
			letters += matrix[z][y]
		}
		allstrings = append(allstrings, letters)
		reverse = reversestring(letters)
		allstrings = append(allstrings, reverse)
		letters = ""
	}

	for y := 0; y < matrixrows; y++ { //		all words by row diagonally increment+1 rows outer loop - Left diagonal half left to right down/up
		w := y
		for z := 0; z < matrixcols; z++ { // javallurmt tmrullavaj ciwiqaulm mluaqiwic obxcebar rabecxbo pilepce ecpelip rorhyr ryhror wcxau uaxcw sllj jlls aao oaa jl lj c
			if w < matrixrows {
				letters += matrix[w][z]
				w++
			}
		}
		allstrings = append(allstrings, letters)
		reverse = reversestring(letters)
		allstrings = append(allstrings, reverse)
		letters = ""
	}

	for y := 1; y <= matrixrows; y++ { //			all words by row diagonally decrement-1 rows outer loop - Left diagonal half left to right up/down
		w := y - 1
		for z := 0; z <= matrixcols; z++ { //  j ce ec oaf fao pimb bmip rbvdl ldvbr wiwocp pcowiw soxakie eikaxos aclispmp pmpsilca jlrclqrgr rgrqlcrlj caxeqerjte etjreqexac
			if w >= 0 {
				//		fmt.Printf("matrix[w][z] = matrix[%v][%v]\n", w, z)
				letters += matrix[w][z]
				w--
			}
		}
		allstrings = append(allstrings, letters)
		reverse = reversestring(letters)
		allstrings = append(allstrings, reverse)
		letters = ""
	}

	for y := 0; y < matrixcols; y++ { //			all words by row diagonally increment+1 rows outer loop - Right diagonal half left to right down/up
		w := y
		for z := 0; z < matrixrows; z++ { // javallurmt tmrullavaj emoseimyp pymiesome fdkqlrgi igrlqkdf bcprhpr rphrpcb lirorc croril pmjus sujmp egsa asge ptm mtp rc cr e
			if w < matrixcols {
				letters += matrix[z][w]
				w++
			}
		}
		allstrings = append(allstrings, letters)
		reverse = reversestring(letters)
		allstrings = append(allstrings, reverse)
		letters = ""
	}

	for z := 0; z < matrixrows; z++ { //                 all words by row diagonally increment+1 rows outer loop - Right diagonal half right to left down/up
		w := z
		for y := matrixcols - 1; y >= 0; y-- { //    etjreqexac caxeqerjte csollehll llhellosc muhiapao oapaihum arrubyj jyburra spmucu ucumps cgrar rargc ryle elyr imr rmi pm mp t

			if w < matrixcols {
				letters += matrix[w][y]
				w++
			}
		}
		allstrings = append(allstrings, letters)
		reverse = reversestring(letters)
		allstrings = append(allstrings, reverse)
		letters = ""
	}

	for z := matrixrows - 1; z >= 0; z-- { //                 all words by row diagonally increment+1 rows outer loop - Right diagonal half right to left up/down
		w := z
		for y := matrixcols - 1; y >= 0; y-- { //    tmrullavaj javallurmt pymiesome emoseimyp igrlqkdf fdkqlrgi rphrpcb bcprhpr croril lirorc sujmp pmjus asge egsa mtp ptm cr rc e]

			if w >= 0 {
				letters += matrix[w][y]
				w--
			}
		}
		allstrings = append(allstrings, letters)
		reverse = reversestring(letters)
		allstrings = append(allstrings, reverse)
		letters = ""
	}

	sort.Strings(allstrings)
	allstrings = removeDuplicates(allstrings)
	fmt.Println(allstrings)

	for _, v := range words {

		for c := 0; c < len(allstrings); c++ {

			if strings.Contains(allstrings[c], string(v)) {
				fmt.Printf("Found %v in row %v\n", string(v),allstrings[c])
			}
		}
	}
}

func reversestring(forwards string) string {
	var reverse string

	for i := len(forwards) - 1; i >= 0; i-- {
		reverse += forwards[i : i+1]
	}
	return reverse
}

func setup(wordsearch []string) {
	for i, v := range wordsearch {
		arr := string(v)
		for x := 0; x < len(arr); x++ {
			matrix[i][x] = arr[x : x+1]
		}
	}
}

func removeDuplicates(allstrings []string) []string {
	encountered := map[string]bool{}
	result := []string{}

	for v := range allstrings {
		if encountered[allstrings[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[allstrings[v]] = true
			// Append to result slice.
			result = append(result, allstrings[v])
		}
	}
	// Return the new slice.
	return result
}
