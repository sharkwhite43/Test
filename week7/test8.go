package main

import "fmt"

func main() {
	panic("Cat")
	meow := recover()
	fmt.Println(meow)

}
