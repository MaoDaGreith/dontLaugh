package main

import "fmt"

// callback -> passing a func as an argument
func main() {
	xx := []int{2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(len(xx))
	fmt.Println(average(sum, xx...))
}

func sum(x ...int) int {
	n := 0
	for _, v := range x {
		n += v
	}
	return n
}

func average(f func(x ...int) int, y ...int) float64 {
	var xi []int
	for _, v := range y {
		xi = append(xi, v)
	}

	total := f(xi...)
	return float64(total) / float64(len(xi))
}
