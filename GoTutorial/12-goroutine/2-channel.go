// 通道（channel）是用来传递数据的一个数据结构。
// 通道可用于两个 goroutine 之间通过传递一个指定类型的值来同步运行和通讯。操作符 <- 用于指定通道的方向，发送或接收。如果未指定方向，则为双向通道。
// 自右向左赋值

package main

import "fmt"

func Sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 将sum发送到通道c
}

func main() {
	s := []int{7, 2, 3, 8, -9, 4, 0}

	c := make(chan int)

	go Sum(s[:len(s)/2], c)

	go Sum(s[len(s)/2:], c)

	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}
