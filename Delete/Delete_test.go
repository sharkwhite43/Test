package main

import "testing"

func TestDelete(t *testing.T) {
	result := Delete(15, 7)

	if result != 11 {
		t.Error("Expected 11 , got :", result)
	}
}
