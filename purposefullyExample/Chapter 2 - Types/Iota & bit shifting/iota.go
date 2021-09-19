package main

import "fmt"

const (
	a = iota
	b
	c
)

/*
Within a constant declaration, the predeclared identifier iota represents successive untyped integer constants.
It is reset to 0 whenever the reserved word const appears in the source.
*/

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)

	// bit shifting does what the name says, shifts bits either to the right >> or to the left <<
	// as the example below suggests shifts 6 by 1 bit to the left makes it into 16
	x := 6
	fmt.Printf("%d\t\t%b\n", x, x)

	y := x << 1
	fmt.Printf("%d\t\t%b", y, y)
}

// for more info on iota: https://docs.google.com/document/d/1ckYpi6hcRkaBUEk975f54oGsHYHu7GhzOk7-nOrkNxo/edit#heading=h.q1bcdl6rllax

// for more info on bit shifting: https://docs.google.com/document/d/1ckYpi6hcRkaBUEk975f54oGsHYHu7GhzOk7-nOrkNxo/edit#heading=h.4mp6vkf9oern
