package main

import "fmt"

func main() {
	var a = "initial"
	fmt.Println(a) // string initial

	var b, c int = 1, 2
	fmt.Println(b, c) // int 1 2

	var d = true
	fmt.Println(d) // boolean true

	var e int      // 未赋值默认初始化我零值 bool: false number: 0 string: ''
	fmt.Println(e) // int 0

	f := "short"   // 简约类型声明只能在函数体内 不可以声明为全局变量
	fmt.Println(f) // string short

	var a1 bool
	var b1 int
	var c1 float32
	var d1 string
	var e1 rune                         // int32
	var f1 byte                         // uint8
	fmt.Println(a1, b1, c1, d1, e1, f1) // false 0 0  0 0
}
