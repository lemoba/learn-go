package main

import "fmt"

// 只写
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// pings 只读 pongs 只写
// <- 紧连参数名为读通道， 反之为写通道
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
