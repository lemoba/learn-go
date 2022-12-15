package main

import "fmt"

func main() {
	var numbers = make([]int, 3, 5)

	// 空切片
	var num []int

	printSlice(num)

	if num == nil {
		fmt.Printf("切片为空\n")
	}
	printSlice(numbers)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
