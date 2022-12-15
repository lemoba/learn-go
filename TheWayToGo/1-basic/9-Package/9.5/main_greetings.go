package main

import (
	"./grettings"
	"fmt"
)

func main() {
	name := "James"
	fmt.Println(grettings.GoodDay(name))
	fmt.Println(grettings.GoodNight(name))

	if grettings.IsAM() {
		fmt.Println("Good morning", name)
	} else if grettings.IsAfernoon() {
		fmt.Println("Good afternoon", name)
	} else if grettings.IsEvening() {
		fmt.Println("Good evening", name)
	} else {
		fmt.Println("Good ngiht", name)
	}
}
