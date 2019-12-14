package main

import "fmt"

func say() {
	fmt.Println("Hi Tan")
}

func main() {
	defer say()
	fmt.Println("Hello World")
}
