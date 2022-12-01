package main

type Interval struct {
	start int
	end   int
}

func main() {
	intr := Interval{0, 3}
	intr2 := Interval{end: 5, start: 1}
	intr3 := Interval{end: 5}

	println(intr.end)
	println(intr2.end)
	println(intr3.end)
}
