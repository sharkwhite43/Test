package main

import "fmt"

func tan(x, y int) (int, int) {
	return y, x
}

func main() {
	x, y := tan(20, 15)
	fmt.Println(x, y)

}
