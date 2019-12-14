package main

import "fmt"

func say() {
	fmt.Println("Hi Tan")
}

func port() {
	fmt.Println("Hello Boat")

}
func main() {
	defer say()
	defer port()

}
