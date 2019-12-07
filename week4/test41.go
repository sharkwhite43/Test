package main

import "fmt"

func main() {
	humes := []string{"Tan", "Nan", "Ter"}
	names := []string{"Baot", "Aum", "meena"}
	names = append(names, humes...)
	fmt.Println(names)

}
