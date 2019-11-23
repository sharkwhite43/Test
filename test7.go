package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Index("Hello World", "world"))
	fmt.Println(strings.LastIndex("Hello World", "World"))
}
