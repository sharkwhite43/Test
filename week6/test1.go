package main

import "fmt"

func Mom() {
	fmt.Println("Hello World")
}

func greet(name string) {
	fmt.Println("Hello ", name)
}

func main() {
	Mom()
	greet("Tan")
}
