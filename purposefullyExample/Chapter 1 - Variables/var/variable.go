package main

import "fmt"

var z int // declare variable of TYPE int, automatically ZERO value is assigned to the variable
// if it was string it would be "" if it was boolean it would be false etc.

func main() {
	// scope levels, either package level of func level
	var y = 23 // can be used ouside as well as inside the func
	fmt.Println(y)
}
