package main

import (
	"fmt"
	"regexp"
)

var p = fmt.Println

func main() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	p(match)

	r, _ := regexp.Compile("p([a-z]+)ch")

	p(r.MatchString("peach"))

	p(r.FindString("peach punch"))

	p("idx:", r.FindStringIndex("peach"))
}
