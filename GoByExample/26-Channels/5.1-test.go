package main

import "fmt"

// 只写
func pin(pings chan<- string, msg string) {
	pings <- msg
}

func pon(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
func main() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	pin(c1, "Im go!")

	pon(c1, c2)
	fmt.Println(<-c2)
}
