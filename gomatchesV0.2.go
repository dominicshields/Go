/* gomatches2.go
My first attempt at a Go program of my own devising, the code is in places intentionally naive for clarity.
A very common file operation is to compare two SORTED files on a shared field or string of bytes
This compares two files that are sorted on the byte positions of interest and writes out 4 files
matches1: matches in file 1 format
matches2: matches in file 2 format
in1not2: records in file 1 not in file 2
in2not1: records in file 2 not in file 1
usage is
gomatches <input file 1> <input file 2> <match position file 1> <match position file 2> <matchlength>
N.B. IF THE FILES ARE NOT SORTED ON THE FIELDS/BYTE POSITION OF THE INPUT MATCH POSITIONS THIS WILL FAIL TO PRODUCE CORRECT OUTPUT FOR OBVIOUS REASONS.
     DUPLICATE "KEYS" ARE NOT TREATED ANY DIFFERENTLY THAN ANY OTHER RECORD.
     INTERESTING OBSERVATION: I WROTE THE IDENTICAL PROGRAM IN BORLAND C 25 YEARS AGO, I STILL HAVE IT, THE C EXECUTABLE IS 65KB, THIS PROGRAM IN GO
     HAS AN EXECUTABLE WHICH IS 2065KB
Dominic Shields 06/03/2018 23:12:11 */

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

var file1cnt int
var file2cnt int
var file1EOF int
var file2EOF int

func main() {
	var infile1buf string
	var infile2buf string

	t0 := time.Now()
	fmt.Println(t0.Local())

	in1not2cnt := 0
	in2not1cnt := 0
	match1cnt := 0
	match2cnt := 0
	matchlen1 := 0
	matchlen2 := 0

	if len(os.Args) < 6 {
		fmt.Printf("Error, usage is %v <input file 1> <input file 2> <match position file 1> <match position file 2> <matchlength>. Only %v parameters input>\n", os.Args[0], len(os.Args)-1)
		os.Exit(1)
	}
	// OPEN FILES
	infile1 := openfile(os.Args[1], "READ")
	defer infile1.Close()

	infile2 := openfile(os.Args[2], "READ")
	defer infile2.Close()

	matches1 := openfile("matches1", "WRITE")
	defer matches1.Close()

	matches2 := openfile("matches2", "WRITE")
	defer matches2.Close()

	in1not2 := openfile("in1not2", "WRITE")
	defer in1not2.Close()

	in2not1 := openfile("in2not1", "WRITE")
	defer in2not1.Close()

	infile1pos, _ := strconv.Atoi(os.Args[3])
	infile2pos, _ := strconv.Atoi(os.Args[4])
	matchlen, _ := strconv.Atoi(os.Args[5])

	file1minlen := infile1pos + matchlen
	file2minlen := infile2pos + matchlen

	// ASSIGN FILE POINTERS TO INPUT FILES - SCANNER DOES NOT WORK WELL IN THIS PROGRAM
	infile1ptr := bufio.NewReader(infile1)
	infile2ptr := bufio.NewReader(infile2)

	// READ INPUT FILES
	infile1buf, _ = readfile1(infile1ptr) // NOT CURRENTLY CHECKING FILE READ ERRORS, I LIKE TO LIVE DANGEROUSLY
	infile2buf, _ = readfile2(infile2ptr)

	// CONTROLLING LOOP HERE
	for file1EOF != 9 && file2EOF != 9 { // A BIT OF A HACK AS GO DOES NOT HAVE THE RICH SET OF FILE OPERATIONS THAT C DOES
		matchlen1 = matchlen
		matchlen2 = matchlen

		// DEAL WITH VARIABLE LENGTH RECORDS WITHOUT GOING OUT OF BOUNDS
		if len(infile1buf) < file1minlen {
			matchlen1 = (len(infile1buf) - infile1pos)
		}
		if len(infile2buf) < file2minlen {
			matchlen2 = (len(infile2buf) - infile2pos)
		}

		if matchlen1 < matchlen2 {
			matchlen2 = matchlen1
		} else if matchlen1 > matchlen2 {
			matchlen1 = matchlen2
		}

		infile1cmp := infile1buf[infile1pos-1 : infile1pos-1+matchlen1]
		infile2cmp := infile2buf[infile2pos-1 : infile2pos-1+matchlen2]

		if infile1cmp > infile2cmp {
			fmt.Fprintf(in2not1, "%v\n", infile2buf)
			in2not1cnt++
			infile2buf, _ = readfile2(infile2ptr)

		} else if infile1cmp < infile2cmp {
			fmt.Fprintf(in1not2, "%v\n", infile1buf)
			in1not2cnt++
			infile1buf, _ = readfile1(infile1ptr)

		} else if infile1cmp == infile2cmp {
			fmt.Fprintf(matches1, "%v\n", infile1buf)
			fmt.Fprintf(matches2, "%v\n", infile2buf)
			match1cnt++
			match2cnt++
			infile1buf, _ = readfile1(infile1ptr)
			infile2buf, _ = readfile2(infile2ptr)
		}
	}

	for file1EOF != 9 {
		fmt.Fprintf(in1not2, "%v\n", infile1buf)
		in1not2cnt++
		infile1buf, _ = readfile1(infile1ptr)
	}

	for file2EOF != 9 {
		fmt.Fprintf(in2not1, "%v\n", infile2buf)
		in2not1cnt++
		infile2buf, _ = readfile2(infile2ptr)
	}

	fmt.Printf("Records read from file 1     = %v\n", file1cnt)
	fmt.Printf("Records read from file 2     = %v\n", file2cnt)
	fmt.Printf("Matching records             = %v\n", match1cnt)
	fmt.Printf("Records in file 1 not file 2 = %v\n", in1not2cnt)
	fmt.Printf("Records in file 2 not file 1 = %v\n", in2not1cnt)
	t1 := time.Now()
	fmt.Println(t1.Local())
	fmt.Printf("Runtime = %v\n", t1.Sub(t0))
}

func openfile(filename string, operation string) *os.File {
	var fp *os.File
	var err error
	if operation == "READ" {
		fp, err = os.Open(filename)
	} else {
		fp, err = os.Create(filename)
	}
	if err != nil {
		fmt.Println("Cannot open file ", filename)
		os.Exit(1)
	}
	return fp
}

func readfile1(infile1ptr *bufio.Reader) (string, error) {
	var infile1buf string
	infile1slice, _, infile1err := infile1ptr.ReadLine()
	n := len(infile1slice)
	infile1buf = string(infile1slice[:n])
	file1cnt++
	if infile1err == io.EOF {
		file1EOF = 9
		return infile1buf, infile1err
	} else if infile1err != nil {
		return infile1buf, infile1err
	}
	return infile1buf, infile1err
}

func readfile2(infile2ptr *bufio.Reader) (string, error) {
	var infile2buf string
	infile2slice, _, infile2err := infile2ptr.ReadLine()
	n := len(infile2slice)
	infile2buf = string(infile2slice[:n])
	file2cnt++
	if infile2err == io.EOF {
		file2EOF = 9
		return infile2buf, infile2err
	} else if infile2err != nil {
		return infile2buf, infile2err
	}
	return infile2buf, infile2err
}
