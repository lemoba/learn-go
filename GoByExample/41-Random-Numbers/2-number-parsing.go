package main

import (
	"fmt"
	"strconv"
)

func main() {
	f, _ := strconv.ParseFloat("12.12", 64) // 64表示解析数的位数
	fmt.Println(f)

	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	k, _ := strconv.Atoi("135") // Atoi 是一个基础的 10 进制整型数转换函数
	fmt.Println(k)

	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
