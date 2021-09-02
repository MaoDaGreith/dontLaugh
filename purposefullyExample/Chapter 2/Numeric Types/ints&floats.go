package main

import (
	"fmt"
)

var x int
var y float64

func main() {
	x = 42      //int value only stores
	y = 34.2323 //float value that can be more precise
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	// difference between int and uint etc.
	// uint8 0 to 255 while int8 -128 up to 127 and so forth with uint16 etc.
	// byte -> uint8, rune -> int32
}

// For more info: https://docs.google.com/document/d/1ckYpi6hcRkaBUEk975f54oGsHYHu7GhzOk7-nOrkNxo/edit#heading=h.ry3lovlkdwv0
