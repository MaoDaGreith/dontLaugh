package main

import "fmt"

func main() {
	xi := []int{21, 22, 23, 24, 25, 26, 27, 28, 48}

	fmt.Println(&xi[0])
	fmt.Println(cap(xi))
	a := &xi
	fmt.Printf("%p\n", a)

	xi = append(xi, 123, 234, 345, 456, 678, 789)
	fmt.Println(&xi[0])
	fmt.Println(cap(xi))
	a = &xi
	fmt.Printf("%p\n", a)
}
