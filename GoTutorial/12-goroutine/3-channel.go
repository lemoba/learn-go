// 通道缓冲区
// 通道可以设置缓冲区，通过 make 的第二个参数指定缓冲区大小：
// ch := make(chan int, 100)
package main

import "fmt"

func main() {
	ch := make(chan int, 2)

	ch <- 1
	ch <- 2

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
