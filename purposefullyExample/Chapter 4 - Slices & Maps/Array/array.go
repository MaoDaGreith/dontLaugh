package main

import "fmt"

func main() {
	s := [5]int{2, 4, 5, 6, 7} // array with values
	var x [5]int               //array without values
	fmt.Println(s)
	x[3] = 22
	fmt.Println(x)
}

// Note: using var will initialize the variable with a zero value... It only reached me in this chapter... I'm hopeless
// https://stackoverflow.com/a/53404567
