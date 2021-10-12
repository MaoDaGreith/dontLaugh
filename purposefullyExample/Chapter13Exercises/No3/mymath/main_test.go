package mymath

import (
	"fmt"
	"sort"
	"testing"
)

func TestCenteredAvg(t *testing.T) {
	type tests struct {
		tes    []int
		result float64
	}

	d := func(xi []int) float64 {
		sort.Ints(xi)
		xi = xi[1:(len(xi) - 1)]

		n := 0
		for _, v := range xi {
			n += v
		}

		f := float64(n) / float64(len(xi))
		return f
	}

	a := []int{5, 6, 2, 8, 12, 3, 35, 33}
	b := []int{12, 3, 6, 20, 11, 21, 31}
	c := []int{10, 2, 4, 13, 12, 23, 35}

	xs := []tests{
		{a, d(a)},
		{b, d(b)},
		{c, d(c)},
	}

	for _, v := range xs {
		if CenteredAvg(v.tes) != v.result {
			t.Error("Expected", v.result, "Got", CenteredAvg(v.tes))
		}
	}
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{2, 3, 4, 5, 6, 7}))
	// Output:
	// 4.5
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{3, 4, 5, 6, 7, 8, 10123})
	}
}
