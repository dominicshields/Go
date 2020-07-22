package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	inputhandle1 := openfile(os.Args[1],"READ")
	defer inputhandle1.Close()
	
	outputhandle1 := openfile(os.Args[2],"WRITE")
	defer outputhandle1.Close()

	fileScanner := bufio.NewScanner(inputhandle1)

	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
	}
	
	fmt.Fprintf(outputhandle1, "Some Output\n")
	
}

func openfile(filename string,operation string)  (*os.File) {
	var fp *os.File
	var err error
	if operation == "READ" {
			fp, err = os.Open(filename)
	} else {
			fp, err = os.Create(filename)
	}
	  if err != nil {
      panic(err)
  }

  return fp
}
