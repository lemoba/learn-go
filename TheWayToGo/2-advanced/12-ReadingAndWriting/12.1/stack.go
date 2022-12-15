package main

import (
	"./stack"
	"fmt"
)

func main() {
	//buf := bufio.NewReader(os.Stdin)
	s := new(stack.Stack)
	fmt.Printf("cap = %t\n", s.IsEmpty())
	for x := 1; x <= 10; x++ {
		s.Push(x)
	}
	fmt.Println(s)
	fmt.Printf("len = %d\n", s.Len())
	fmt.Printf("cap = %d\n", s.Cap())
	fmt.Printf("is empty = %t\n", s.IsEmpty())
	s.Pop()
	fmt.Println(s)

	res, err := s.Top()
	if err == nil {
		fmt.Println(res)
	}
}
