package main

import "fmt"

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	PrintSlice(numbers)

	// 打印原始切片
	fmt.Println("numbers == ", numbers)

	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numbers[1:4] == ", numbers[1:4]) // 1 2 3

	/* 默认下限为 0*/
	fmt.Println("numbers[:3] ==", numbers[:3]) // 0 1 2

	/* 默认上限为 len(s)*/
	fmt.Println("numbers[4:] == ", numbers[4:]) // 4 5 6 7 8 9

	num1 := make([]int, 0, 5)
	PrintSlice(num1)

	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	num2 := numbers[:2]
	PrintSlice(num2)

	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	num3 := numbers[2:5]
	PrintSlice(num3)

}

func PrintSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
