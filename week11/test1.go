package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	reader := strings.NewReader("Hello World")
	p := make([]byte, 3)
	for {
		n, err := reader.Read(p)
		if err == io.EOF {
			break
		}
		fmt.Println(string(p[:n]))
	}

}
