package main

import "fmt"

type number struct {
	f float32
}

type nr number // alias type

func main() {
	a := number{5.1}
	b := nr{5.2}
	var c = number(b)
	fmt.Println(a, b, c)
}
