package main

import (
	"fmt"
	"math"
)

/*
const 声明常量
Go支持字符、字符串、布尔和数值常量
*/

const s string = "constant"

func main() {

	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
