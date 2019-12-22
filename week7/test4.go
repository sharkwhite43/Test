package main

import "fmt"

func Toomorro() {
	morning := recover()
	fmt.Println(morning)
}
func toyou() {
	ting := recover()
	fmt.Println(ting)
}
func main() {
	defer Toomorro()
	