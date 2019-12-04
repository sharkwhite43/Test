package main

import "testing"

func TestMultiply(t *testing.T) {
	result := Multiply(4, 3)

	if result != 12 {
		t.Error("Expected 12 , got :", result)
	}

}
