package main

import "testing"

func TestMySum(t *testing.T) {

	type swaggies struct {
		bucks   []int
		fingers int
	}

	drip := []swaggies{
		{[]int{32, 35, 2}, 69},
		{[]int{11, 31}, 42},
		{[]int{10, 12, -5}, 17},
		{[]int{200, -30, -1}, 169},
	}

	for _, v := range drip {
		if muscles(v.bucks...) != v.fingers {
			t.Error("Expected", v.fingers, "Got", muscles(v.bucks...))
		}
	}
}
