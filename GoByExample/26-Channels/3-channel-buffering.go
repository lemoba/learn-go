package main

import "fmt"

func main() {
	msg := make(chan string, 2)

	msg <- "buffered"
	msg <- "channel1"

	fmt.Println(<-msg)
	fmt.Println(<-msg)
}
