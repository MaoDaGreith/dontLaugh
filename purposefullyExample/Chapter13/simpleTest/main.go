package main

import "fmt"

func main() {
	fmt.Println("2 + 7", mySum(2, 7))
}

func mySum(a, b int) int {
	sum := a + b
	return sum
}
