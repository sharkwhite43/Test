package main

import "fmt"

func main() {
	i := 0
	for {
		fmt.Println(i)
		i = i + 5*2
		if i >= 100 {
			i = i + 10 //++
			fmt.Println(i)
			break

		}
	}
}
