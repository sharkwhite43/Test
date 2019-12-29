package main

import "fmt"

type I interface{}

func desc(i I) {
	fmt.Println("%v , %T \n", i, i)
}

func main() {
	var i T
	i = 3.1459265359
	desc(i)
	i = "Hello World"

}
