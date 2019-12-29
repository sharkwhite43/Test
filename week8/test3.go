package main

import "fmt"

type student struct {
	name  string
	age   int
	email string
}

func main() {
	var a student
	a.name = "Tan"
	a.age = 19
	a.email = "TanTawan23@gmail.com"

	b := student{"Tle", 19, "Tle@mail.com"}

	c := student{name: "Ter", email: "Ter@mail.com"}

	d := student{age: 22}

	fmt.Println(a)
	fmt.Println(b)

}
