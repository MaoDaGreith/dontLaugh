package main

import "fmt"

func main() {
	s := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for _, v := range s {
		fmt.Println(v)
	}
	fmt.Printf("\n%T", s)
}
