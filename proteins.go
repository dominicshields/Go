package main

import (
	"fmt"
	//	"math"
	//	"sort"
	//	"strconv"
)

func main() {
	const RNA = "AUGUUUUCUUGUUUAUAA"
	var codes map[string]string
	var codon string
	var protein []string

	codes = map[string]string{
		"AUG": "Methionine",
		"UUU": "Phenylalanine",
		"UUC": "Phenylalanine",
		"UUA": "Leucine",
		"UUG": "Leucine",
		"UCU": "Serine",
		"UCC": "Serine",
		"UCA": "Serine",
		"UCG": "Serine",
		"UAU": "Tyrosine",
		"UAC": "Tyrosine",
		"UGU": "Cysteine",
		"UGC": "Cysteine",
		"UGG": "Tryptophan",
		"UAA": "STOP",
		"UAG": "STOP",
		"UGA": "STOP",
	}

	for x := 0; x < len(RNA); x = x + 3 {
		codon = RNA[x : x+3]
		protein = append(protein, codes[codon])
	}
	fmt.Printf("Protein : %v\n", protein)
}
