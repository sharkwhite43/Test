package main

import "fmt"

func main() {
	humans := []string{"Tan", "Nan", "Ter"}
	names := []string{"Baot", "Aum", "meena"}
	names = append(names, humans...)
	fmt.Println(names)

}
