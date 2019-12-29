package main

import "fmt"

type student struct {
	name  string
	age   int
	email string
}

func main() {
	std := student{name: "Ball"}
	p := &std
	(*p).age = 18
	p.email = "Ball@gmail.com"

	fmt.Println(std)

}
