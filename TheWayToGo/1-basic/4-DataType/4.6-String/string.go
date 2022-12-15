package main

import "fmt"

/**
 *
 */
func main() {
	str := `This is a row string \n`
	str1 := "Beginning of the string" + "second part of the string" // + 必须放在第一行
	str2 := "hel" + "lo, "
	str2 += "world!"
	fmt.Println(str)
	fmt.Println(str1)
	fmt.Println(str2)
}
