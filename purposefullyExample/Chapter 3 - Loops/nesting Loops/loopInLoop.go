package main

import "fmt"

func main() {
	// A stpry about Timmy watching Shrek in a loop statement
	for i := 0; i <= 10; i++ {
		fmt.Printf("Timmy wasted %d hours\n", i)
		for j := 1; j <= 3; j++ {
			fmt.Printf("\t\t\tin Shrek %d\n", j)
		}
	}
}
