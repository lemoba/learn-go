package main

import (
	"./stack"
)

func main() {
	stack := new(stack.Stack)

	for i := 0; i < 10; i++ {
		stack.Push(i)
	}
}
