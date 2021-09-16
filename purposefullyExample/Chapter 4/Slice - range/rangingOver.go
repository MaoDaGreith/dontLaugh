package main

import "fmt"

func main() {
	// COMPOSITE LITERAL
	s := []int{1, 3, 5, 7, 8} // slices position start from 0
	fmt.Println(s[0])
	fmt.Println(s[1])
	fmt.Println(s[2])
	fmt.Println(s[3])
	fmt.Println(s[4])

	for i, v := range s {
		fmt.Println(i, v) // i is for index(position in slice) and v is for the corresponding value from the slice
	}
}
