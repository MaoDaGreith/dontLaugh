package main

import "fmt"

func main() {
	s := []int{12, 13, 14, 15, 16, 17, 18}
	fmt.Println(foo(s...))

	fmt.Println(bar(s))
}

func foo(t ...int) int {
	var total int
	for _, v := range t {
		total += v
	}
	return total
}

func bar(x []int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
