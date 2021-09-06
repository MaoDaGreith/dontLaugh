package main

import "fmt"

func main() {
	a := 42
	fmt.Printf("%d\t%b\t%#X\n", a, a, a)
	b := 42 << 1
	fmt.Printf("%d\t%b\t%#X", b, b, b)
}
