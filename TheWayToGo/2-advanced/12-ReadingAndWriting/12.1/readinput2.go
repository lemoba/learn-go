package main

import (
	"bufio"
	"fmt"
	"os"
)

var inputReader *bufio.Reader
var _input string
var err error

func main() {
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input：")
	_input, err = inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("The input was: %s\n", _input)
	}
}
