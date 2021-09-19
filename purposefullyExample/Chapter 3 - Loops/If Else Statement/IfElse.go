package main

import "fmt"

func main() {
	n := "number"

	if n == "double" {
		fmt.Println("Not a double")
	} else if n == "syke" {
		fmt.Println("different than Double")
	} else {
		fmt.Println("none of the above options")
	}
}
