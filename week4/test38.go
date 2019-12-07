package main

import "fmt"

func main() {
	alphabets := [4]string{"A", "B", "C", "D"}
	x := alphabets[:]
	y := alphabets[:3]
	z := alphabets[2:]
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

}
