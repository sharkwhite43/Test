package main

import "fmt"

func mian() {
	var p *int
	i := 42
	fmt.Println("value", i)
	p = &i
	*p = 3
	fmt.Println("value", i)
}
