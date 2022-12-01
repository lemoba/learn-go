package main

func vals() (int, int) {
	return 3, 7
}

func main() {
	a, b := vals()
	println(a) // 3
	println(b) // 7

	_, c := vals()
	println(c) // 7
}
