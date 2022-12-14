package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)
	fmt.Println("len:", len(s))

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]            // 前闭后开
	fmt.Println("sl1:", l) // c d e

	l = s[2:]
	fmt.Println("sl2:", l) // c d e f

	l = s[:5]
	fmt.Println("sl3:", l) // a b c d e

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t) // g h i

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
