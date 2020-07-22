/* grep.go
The grep command takes three arguments:
The pattern used to match lines in a file.
Zero or more flags to customize the matching behavior.
One or more files in which to search for matching lines.
the grep command should also support the following flags:

-i Match line using a case-insensitive comparison.
-l Print only the names of files that contain at least one matching line.
-n Print the line numbers of each matching line.
-v Invert the program -- collect all lines that fail to match the pattern.
-x Only match entire lines, instead of lines that contain a match.
Example call without flags
grep "hello" input.txt
Example call with flags
grep -n "hello" input.txt

*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
//	"time"
)

var file1cnt int
var file1EOF int
var linecnt int64

func main() {
	var infile1 *os.File
	var infile1buf string
	var grepstr string
	var inputfile string
	iflag := 0
	lflag := 0
	nflag := 0
	vflag := 0
	xflag := 0
	linecnt = 0

//	t0 := time.Now()
//	fmt.Println(t0.Local())

	if len(os.Args) < 3 {
		fmt.Printf("Error, usage is %v [-nlivx] <search string> <input file>\n", os.Args[0])
		os.Exit(1)
	}

	params := os.Args[1]

	if params[0:1] == "-" {
		for _, y := range params {
			switch string(y) {
			case "i":
				iflag++
			case "l":
				lflag++
			case "n":
				nflag++
			case "v":
				vflag++
			case "x":
				xflag++
			}
		}
		infile1 = openfile(os.Args[3], "READ")
		inputfile = os.Args[3]
		grepstr = os.Args[2]
	} else {
		infile1 = openfile(os.Args[2], "READ")
		inputfile = os.Args[2]
		grepstr = os.Args[1]
	}
	defer infile1.Close()

	infile1ptr := bufio.NewReader(infile1)
	infile1buf, _ = readfile(infile1ptr, 1)

	/*-i Match line using a case-insensitive comparison.
	  -l Print only the names of files that contain at least one matching line.
	  -n Print the line numbers of each matching line.
	  -v Invert the program -- collect all lines that fail to match the pattern.
	  -x Only match entire lines, instead of lines that contain a match. */

	for file1EOF != 9 {

		if iflag == 0 && lflag == 0 && nflag == 0 && vflag == 0 && xflag == 0 {     // no flags
			if strings.Contains(infile1buf, grepstr) {
				fmt.Printf("%v\n", infile1buf)
			}
		}

		if iflag != 0 && lflag == 0 && nflag == 0 && vflag == 0 && xflag == 0 {				// case insensitive flag
			if strings.Contains(strings.ToUpper(infile1buf), strings.ToUpper(grepstr)) {
				fmt.Printf("%v\n", infile1buf)
			}
		}
		
		if iflag != 0 && lflag == 0 && nflag != 0 && vflag == 0 && xflag == 0 {				// case insensitive flag + line numbering flag
			if strings.Contains(strings.ToUpper(infile1buf), strings.ToUpper(grepstr)) {
				fmt.Printf("%v:  %v\n", linecnt, infile1buf)
			}
		}
		
		if iflag != 0 && lflag != 0 && nflag == 0 && vflag == 0 && xflag == 0 {				// case insensitive flag + name of file only flag
			if strings.Contains(strings.ToUpper(infile1buf), strings.ToUpper(grepstr)) {
				fmt.Printf("%v found in %v\n", grepstr,inputfile)
				break
			}
		}
		
		if iflag != 0 && lflag != 0 && nflag == 0 && vflag != 0 && xflag == 0 {				// case insensitive flag + name of file only flag + do not match flag
			if !strings.Contains(strings.ToUpper(infile1buf), strings.ToUpper(grepstr)) {
				fmt.Printf("%v not found in %v\n", grepstr,inputfile)
				break
			}
		}
		
		if iflag != 0 && lflag != 0 && nflag == 0 && vflag != 0 && xflag != 0 {				// case insensitive flag + name of file only flag + do not match flag + match entire line flag
			if strings.ToUpper(infile1buf) != strings.ToUpper(grepstr) { 
				fmt.Printf("%v not found in %v\n", grepstr,inputfile)
				break
			}
		}
		
		if iflag == 0 && lflag != 0 && nflag == 0 && vflag == 0 && xflag == 0 {	// name of file only flag
			if strings.Contains(infile1buf, grepstr) {
				fmt.Printf("%v found in %v\n", grepstr,inputfile)
				break
			}
		}

		if iflag == 0 && lflag == 0 && nflag != 0 && vflag == 0 && xflag == 0 {	// line numbering flag
			if strings.Contains(infile1buf, grepstr) {
				fmt.Printf("%v:  %v\n", linecnt, infile1buf)
			}
		}

		if iflag == 0 && lflag == 0 && nflag == 0 && vflag != 0 && xflag == 0 {	// lines that do not match flag
			if !strings.Contains(infile1buf, grepstr) {
				fmt.Printf("%v\n", infile1buf)
			}
		}

		if iflag == 0 && lflag == 0 && nflag == 0 && vflag == 0 && xflag != 0 {	// match entire line flag
			if infile1buf == grepstr {
				fmt.Printf("%v\n", infile1buf)
			}
		}

		infile1buf, _ = readfile(infile1ptr, 1)

	}  // for loop

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

func readfile(infileptr *bufio.Reader, filenum int) (string, error) {
	var infilebuf string
	infileslice, _, infileerr := infileptr.ReadLine()
	linecnt++

	n := len(infileslice)
	infilebuf = string(infileslice[:n])
	if infileerr == io.EOF {
		if filenum == 1 {
			file1EOF = 9
		}
		return infilebuf, infileerr
	} else if infileerr != nil {
		return infilebuf, infileerr
	}
	return infilebuf, infileerr
}
