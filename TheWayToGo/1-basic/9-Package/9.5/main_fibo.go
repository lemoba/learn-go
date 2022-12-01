package main

import (
	"./fibo"
	"fmt"
)

var nextFibo int
var op string

func main() {
	op = "+"
	calls()
	fmt.Println("Change of operation from + to *")
	nextFibo = 0
	op = "*"
	calls()
}

func calls() {
	next()
	fmt.Println("...")
	next()
	fmt.Println("...")
	next()
	fmt.Println("...")
	next()
}

func next() {
	result := 0
	nextFibo++
	result = fibo.Fibonacci(op, nextFibo)
	fmt.Printf("fibonacci(%d) is: %d\n", nextFibo, result)
}
