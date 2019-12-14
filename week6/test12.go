package main

import "fmt"

func compyter(fn func(int, int) int) int {
	return fn(3, 4)
}

func main() {
	sum := func(x, y int) int {
		return x + y
	}
	subtract := func(x, y int) int {
		return x - y
	}
	x := computer(sum)
	y := computer(suntract)
	fmt.Println(x)
}
