package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	inputFile, inputError := os.Open("E:\\workspace\\go\\learn-go\\2-advanced\\12-ReadingAndWriting\\12.2\\input.dat")
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got access to it?\n")
		return
	}

	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		fmt.Printf("The input was: %s", inputString)
		if readerError == io.EOF {
			return
		}
	}
}
