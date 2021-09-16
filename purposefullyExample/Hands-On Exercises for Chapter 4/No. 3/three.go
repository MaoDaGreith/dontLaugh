package main

import "fmt"

func main() {
	s := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	a := append(s[:5])
	b := append(s[5:])
	c := append(s[2:7])
	d := append(s[1:6])
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
