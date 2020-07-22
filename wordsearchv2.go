package main

import (
	"fmt"
	"sort"
	"strings"
)

const matrixrows = 10
const matrixcols = 10

type allstring struct {
	lets string
	posy int
	posz int
}

var matrix [matrixrows][matrixcols]string

func main() {
	var wordsearch = []string{"jefblpepre", "camdcimgtc", "oivokprjsm", "pbwasqroua", "rixilelhrs", "wolcqlirpc", "screeaumgr", "alxhpburyi", "jalaycalmp", "clojurermt"}
	var words = []string{"clojure", "elixir", "ecmascript", "rust", "java", "lua", "lisp", "ruby"}
	allstrings := []allstring{}
	var letters string
	var reverse string
	setup(wordsearch)

	for y := 0; y < matrixrows; y++ { //			all words by row
		for z := 0; z < matrixcols; z++ { // jefblpepre erpeplbfej camdcimgtc ctgmicdmac oivokprjsm msjrpkovio pbwasqroua auorqsawbp rixilelhrs srhlelixir wolcqlirpc cprilqclow screeaumgr rgmuaeercs alxhpburyi iyrubphxla jalaycalmp pmlacyalaj clojurermt tmrerujolc
			letters += matrix[y][z]
		}
		allstrings = append(allstrings, allstring{letters, y, 0})
		reverse = reversestring(letters)
		allstrings = append(allstrings, allstring{reverse, y, 9})
		letters = ""
	}

	for y := 0; y < matrixcols; y++ { //			all words by column
		for z := 0; z < matrixrows; z++ { // jcoprwsajc cjaswrpocj eaibioclal lalcoibiae fmvwxlrxlo olxrlxwvmf bdoaicehaj jaheciaodb lckslqepyu uypeqlskcl pipqelabcr rcbaleqpip emrrliuuae eauuilrrme pgjohrmrlr rlrmrhojgp rtsurpgymm mmygprustr ecmascript tpircsamce
			letters += matrix[z][y]
		}
		allstrings = append(allstrings, allstring{letters, 0, y})
		reverse = reversestring(letters)
		allstrings = append(allstrings, allstring{reverse, 9, y})
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
		allstrings = append(allstrings, allstring{letters, y, 9})
		reverse = reversestring(letters)
		allstrings = append(allstrings, allstring{reverse, y, 0})
		letters = ""
	}

	for y := 1; y <= matrixrows; y++ { //			all words by row diagonally decrement-1 rows outer loop - Left diagonal half left to right up/down
		w := y - 1
		for z := 0; z <= matrixcols; z++ { //  j ce ec oaf fao pimb bmip rbvdl ldvbr wiwocp pcowiw soxakie eikaxos aclispmp pmpsilca jlrclqrgr rgrqlcrlj caxeqerjte etjreqexac
			if w >= 0 {
				letters += matrix[w][z]
				w--
			}
		}
		allstrings = append(allstrings, allstring{letters, y - 1, 0})
		reverse = reversestring(letters)
		allstrings = append(allstrings, allstring{reverse, y - 1, 9})
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
		allstrings = append(allstrings, allstring{letters, 0, y})
		reverse = reversestring(letters)
		allstrings = append(allstrings, allstring{reverse, 9, y})
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
		allstrings = append(allstrings, allstring{letters, z, 9})
		reverse = reversestring(letters)
		allstrings = append(allstrings, allstring{reverse, z, 0})
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
		allstrings = append(allstrings, allstring{letters, z, 0})
		reverse = reversestring(letters)
		allstrings = append(allstrings, allstring{reverse, z, 9})
		letters = ""
	}

	sort.Slice(allstrings, func(i, j int) bool {
		return allstrings[i].lets < allstrings[j].lets
	})

	allstrings = removeDuplicates(allstrings)
//	fmt.Println(allstrings)

	for _, v := range words {

		for c := 0; c < len(allstrings); c++ {

			if strings.Contains(allstrings[c].lets, string(v)) {
				fmt.Printf("Found %v in row %v. Start Position Row = %v, Column = %v\n", string(v), allstrings[c].lets, allstrings[c].posy+1, allstrings[c].posz+1)
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

func removeDuplicates(allstrings []allstring) []allstring {
	encountered := map[string]bool{}
	result := []allstring{}

	for v := range allstrings {
		if encountered[allstrings[v].lets] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[allstrings[v].lets] = true
			// Append to result slice.
			result = append(result, allstring{allstrings[v].lets, allstrings[v].posy, allstrings[v].posz})
		}
	}
	// Return the new slice.
	return result
}
