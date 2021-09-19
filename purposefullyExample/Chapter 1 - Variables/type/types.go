package main

import "fmt"

var y = 42                                  // y is of value int
var z = "Shrek 3 is the best movie ever!!!" // z is of value string
var x string                                // x is of value string
var a string = `he said the following 		
"Oh well"` //using `` instead of "" will store the double quotes and new line in a string if you need to

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y) // %T to display the type of the variable
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
