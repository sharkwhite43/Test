package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Split("Hello World", "world"))
	fmt.Println(strings.Split("Hello World", "World"))
}