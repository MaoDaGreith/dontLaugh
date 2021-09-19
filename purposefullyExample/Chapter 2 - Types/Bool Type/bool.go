package main

import "fmt"

var x bool

func main() {
	fmt.Println(x) //printing the value of x which has assigned the "zero value" false
	a := 7
	b := 42
	fmt.Println(a == b) // == is for comparison if it's equal (false)
	fmt.Println(a != b) // != is for comparison if it's not equal (true)
	fmt.Println(a >= b) // >= is for comparison if it's larger than (false)
	fmt.Println(a <= b) // <= is for comparison if it's smaller than (true)
}

// For more info: https://docs.google.com/document/d/1ckYpi6hcRkaBUEk975f54oGsHYHu7GhzOk7-nOrkNxo/edit#heading=h.2ezcntjucx3x
