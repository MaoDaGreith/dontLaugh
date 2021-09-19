package main

import "fmt"

func main() {
	if i := 42; i == 32 {
		fmt.Println("It's true")
	} // Non-ordinary if statement, when i variable is used inside the if statement

	if 43 == 43 {
		fmt.Println("not True")
	} // If statement that compares 2 numbers, and it prints based on evaluating the expression as true
}
