package main

import "fmt"

func main() {
	add := func(x, y int) int {
		return x + y
	}
	x := add(156, 214)
	fmt.Println(x)
}
