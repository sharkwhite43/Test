package main

import "fmt"

func main() {
	alphabets := [2][3]string{{"a", "b", "c"}}
	fmt.Println(alphabets)
	fmt.Println(alphabets[0][1])
	numbers := [2] [3] [2] int{
		{
			{1 ,2} ,
			{10 , 20} ,
			{100 , 200} , 
		} , 
		{
			{8 , 9}
			{80 , 90},
			{800 , 900},
		},
	}

}
