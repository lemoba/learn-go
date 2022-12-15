package main

import "fmt"

func main() {
	var i uint64
	for i = 0; i < 10; i++ {
		fmt.Printf("%d\n", Fibonacci(i))
	}
}

func Fibonacci(n uint64) (result uint64) {
	if n < 2 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
