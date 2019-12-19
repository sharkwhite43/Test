package main

import "fmt"

func handlePanuc() {
	text := recover()
	fmt.Println(text)

}
func main() {
	defer handlePanic()

}
