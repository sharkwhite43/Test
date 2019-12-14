package main

import "fmt"

func subtract(x, y int) int {
	ans := x + y
	return ans
}

func main() {
	x := subtract(10, 5)
	fmt.Println(x)
}
