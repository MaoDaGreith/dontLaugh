package main

import "fmt"

var y = 42

func main() {
	fmt.Printf("%T\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%x\n", y)
	fmt.Printf("%#x\n", y)

	fmt.Printf("%T\t%b\t%x\t%#x\n", y, y, y, y)
	x := fmt.Sprintf("%T\t%b\t%x\t%#x", y, y, y, y) // Sprint is used to print to a string
	fmt.Println(x)

	//Fprint is used to print into a file (somewhere on the server or etc.)
}
