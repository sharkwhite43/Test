package Test25

import "testing"

run test | debug test
func TestAverage(t *testing.T)  {
	v :=Average (3,5)
	if v != 4 {
		t.Error("Average of 3,5 is 4, not" , v)
	}
	

}