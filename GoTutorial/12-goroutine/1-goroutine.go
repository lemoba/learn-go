/*
goroutine 是轻量级线程，goroutine 的调度是由 Golang 运行时进行管理的
Go 允许使用 go 语句开启一个新的运行期线程， 即 goroutine，以一个不同的、新创建的 goroutine 来执行一个函数。
同一个程序中的所有 goroutine 共享同一个地址空间
*/
package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
func main() {
	go say("world")
	say("hello")
}
