package main

import "fmt"

type I interface{}

func desc(i I) {
	fmt.Println("%v , %T \n", i, i)

}
