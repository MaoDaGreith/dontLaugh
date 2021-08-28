package main

import "fmt"

func main() {
	// scope levels, either package level of func level
	var y = 23 // can be used ouside as well as inside the func
	fmt.Println(y)
}
