package main

import "fmt"

func main() {
	// `var`声明1个或者多个变量
	var a string = "ch1 world"
	fmt.Println(a)

	// 一次性声明多个变量
	var b, c int = 1, 2
	fmt.Println(b, c)

	// 自动推断已经初始化的变量类型
	var d = true
	fmt.Println(d)

	// 声明后却没有给出对于的初始值，变量将会被初始化为零值
	var e int
	fmt.Println(e)

	// `:=`声明并初始化 只能在函数体内使用，不可在函数外使用
	f := "ch1"
	fmt.Println(f)
}
