package main

import "fmt"

// 递归函数，就是在运行中自己调用自己
/*
function recursion {
	recursion();
}

func main() {
	recursion();
}


*/
func main() {
	var i int = 4
	fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(uint64(i)))
}

func Factorial(n uint64) uint64 {
	if n > 1 {
		result := n * Factorial(n-1)
		return result
	}
	return 1
}
