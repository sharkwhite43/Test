package main

import "fmt"

type student struct {
	name  string
	age   int
	email string
}

func main() {
	var std [10]student
	std[0] = student{"Tan", 19, "Tantawanais@gmail.com"}

	fmt.Println(std[0])
	fmt.Println(std[0].name)
	fmt.Println(std[0].age)

}
