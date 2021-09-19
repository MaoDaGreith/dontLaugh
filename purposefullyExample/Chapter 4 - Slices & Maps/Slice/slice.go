package main

import "fmt"

func main() {
	// COMPOSITE LITERAL
	s := []int{1, 3, 5, 7, 8}
	fmt.Println(len(s))
	fmt.Println(cap(s))
}

// a SLICE allows you to group together VALUES of the same TYPE
