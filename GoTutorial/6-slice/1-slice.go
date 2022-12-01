package main

import "fmt"

// 切片："动态数组"
// 长度不固定
func main() {
	// 1.定义
	// var indentifier []type
	// var indentifier []type = make([]type, len]
	//slice := make([]int, 10, 20)
	//
	//fmt.Printf("len: %d\ncap: %d\n", len(slice), cap(slice))

	// 2.初始化
	s := []int{1, 2, 3, 4, 5}
	fmt.Printf("len: %d\ncap: %d\n", len(s), cap(s))
}
