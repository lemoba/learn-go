package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "This is an example of a string"
	fmt.Printf("T/F? Does the string \"%s\" hava prefix %s? ", str, "Th")
	fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))
}

func HasPrefix(str string, prefix string) bool {
	return len(str) > len(prefix) && str[0:len(prefix)] == prefix
}
