package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("this is false")
	case 2 != 2:
		fmt.Println("equals 1")
		//fallthrough
	case 2 != 2:
		fmt.Println("Not Equals")
	default:
		fmt.Println("Default value")
	}

	n := 42
	switch n {
	case 43, 41, 40:
		fmt.Println("this is false")
	case 55:
		fmt.Println("equals 1")
		//fallthrough
	case 42:
		fmt.Println("Answer to all the questions")
	default:
		fmt.Println("Default value")
	}
}
