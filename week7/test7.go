package main

import "fmt"

func main() {
	panic("Team")
	Team := recover()
	fmt.Println(Team)

}
