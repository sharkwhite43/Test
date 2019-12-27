package main

import "fmt"

func main() {
	var studentName [10]string
	var studentAge [10]int
	var studentEmail [10]string

	studentName[0] = "Tan"
	studentAge[0] = 19
	studentEmail[0] = "Tantawan.ch"

	fmt.Println(studentName[0], studentAge[0], studentEmail[0])

}
