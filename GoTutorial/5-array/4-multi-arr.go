package main

import "fmt"

/*
多维数组
*/

func multiArr() {
	values := [][]int{}

	row1 := []int{1, 2, 3}
	row2 := []int{4, 5, 6}

	values = append(values, row1)
	values = append(values, row2)

	fmt.Println("Row 1")
	fmt.Println(values[0])
	fmt.Println("Row 2")
	fmt.Println(values[1])

	fmt.Printf("第一个元素：")
	fmt.Println(values[0][0])

	fmt.Println("遍历数组：")

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(values[i][j])
		}
	}
}

// 初始化二维数组
func multiArrInit() {
	// 创建二维数组
	sites := [2][2]string{}

	// 向二维数组添加元素
	sites[0][0] = "Google"
	sites[0][1] = "Runoob"
	sites[1][0] = "Taobao"
	sites[1][1] = "Weibo"
	//
	//fmt.Println(sites)
	PrintArr(sites)

}

func PrintArr(a [2][2]string) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			fmt.Printf("a[%d][%d] = %s\n", i, j, a[i][j])
		}
	}
}

func Animal() {
	animals := [][]string{}

	row1 := []string{"fish", "shark", "eel"}
	row2 := []string{"bird"}
	row3 := []string{"lizard", "salamander"}

	animals = append(animals, row1)
	animals = append(animals, row2)
	animals = append(animals, row3)

	for i := range animals {
		fmt.Printf("Row: %v\n", i)
		fmt.Println(animals[i])
	}
}
func main() {
	multiArr()
	multiArrInit()
	Animal()
}
