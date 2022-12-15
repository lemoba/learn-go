package pack1

import "fmt"

var Pack1Int int = 42
var pack1Float = 3.14

func init() {
	fmt.Printf("init\n")
}

func RetrunStr() string {
	return "Hello main!"
}
