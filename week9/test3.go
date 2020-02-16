package main

import "io"

func main() {
	reader := string.NewReader("HelloWorld")
	p := make([]byte, 3)
	for {
		n, err := reader.Read(p)
		if err == io.EOF {
			break

		}
	}

}
