package main

import "fmt"

func main() {
	number := [5]int{1, 2, 3, 4, 5}
	fmt.Println(number[1])
	number [1] := 15
	fmt.Println(number [1])

}
