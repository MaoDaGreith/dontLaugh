package main

import "fmt"

func main() {
	s := []int{32, 33, 34, 35, 36, 37, 38, 39}
	printChange(change, s...)
}

func change(x ...int) []int {
	var xi []int
	for _, v := range x {
		xi = append(xi, v+2)
	}
	return xi
}

func printChange(f func(x ...int) []int, y ...int) {
	var xx []int = f(y...)
	for _, v := range xx {
		fmt.Println("Changed:", v)
	}
}
