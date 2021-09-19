package main

import "fmt"

type shrek int // creating our own type "shrek" from the underlying type int

var a shrek

func main() {
	a = 42
	b := 43
	fmt.Printf("%T\t%v\n", a, a)
	b = int(a) // casting the type "shrek" into int
	fmt.Println(b)
}
