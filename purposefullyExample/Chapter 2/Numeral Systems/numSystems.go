package main

import "fmt"

func main() {
	s := "V"
	fmt.Println(s)

	b := []byte(s)
	fmt.Println(b)

	n := b[0]
	fmt.Println(n)
	fmt.Printf("%T\n", n)  // %T to display the type of the variable (constant)
	fmt.Printf("%b\n", n)  // to display binary values use %b
	fmt.Printf("%#x\n", n) // hex values can either be %x for lowercase or %X for uppercase also if you want the "0x" use %#X
}

// for more info: https://docs.google.com/document/d/1ckYpi6hcRkaBUEk975f54oGsHYHu7GhzOk7-nOrkNxo/edit#heading=h.vr5aigivfqc2
