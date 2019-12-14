package main

import "fmt"

func suntract(number1, number2 int) {
	number1 = number1 - 1
	number2 = number2 - 1

}

func main() {
	x := 10
	y := 15
	suntract(x, y)
	fmt.Println(x, y)

}
