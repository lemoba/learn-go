package main

import "fmt"

func main() {
	fmt.Println("go" + "lang") // go lang

	fmt.Println("1+1 = ", 1+1)                      // 2
	fmt.Printf("9.7524 / 2.1 = %.2f\n", 9.7524/2.1) // 四舍五入

	fmt.Println(true && false) // false
	fmt.Println(true || false) // true
	fmt.Println(!true)         // false
}
