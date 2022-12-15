package main

import "fmt"

/*
数组是具有相同唯一类型的一组已编号且长度固定的数据项序列
*/
func main() {
	// 声明
	// var balance [5]float32

	// 初始化

	// 方式1
	// var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	// 长度不确定
	// var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

	// 方式2
	// balance := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	// 长度不确定
	// balance := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

	// 方式3
	// 如何设置了数组长度 可以通过指定下标来初始化元素
	balance := [5]float32{1: 2.0, 3: 7.0}
	fmt.Print(balance)
}
