package main

import "fmt"

var y string
var z int

func main() {
	fmt.Println("The zero value of y is: ", y, "finish")
	fmt.Printf("%T\n", y)

	y = "Shrek mlg video"
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
}
