package main

import "fmt"

func main() {
	fmt.Println("2 + 7 =", muscles(2, 7))
}

func muscles(a ...int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}
