package main

import "fmt"

func main() {
	if 12 != 12 {
		fmt.Println("Not in action")
	} else if 12 == 11 {
		fmt.Println("if in Action")
	} else {
		fmt.Println("Shrek is Love, Shrek is Life")
	}
}
