package main

import (
	"./pack1"
	"fmt"
)

func main() {
	var test1 string
	test1 = pack1.RetrunStr()
	fmt.Printf("RetrunStr from package1: %s\n", test1)
	fmt.Printf("Integer from package1: %d\n", pack1.Pack1Int)
	//fmt.Printf("Float from package1: %f\n", pack1.pack1Float)  // 小写表示praivte不可导出
}
