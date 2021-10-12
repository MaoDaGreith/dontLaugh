package main

import "testing"

func TestMySum(t *testing.T) {
	if mySum(2, 5) != 2+5 {
		t.Error("Expected", 7, "Got", mySum(2, 5))
	}
}
