package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ { //for init; condition; post
		if i%2 == 0 { // Verifies if i is an even number
			fmt.Println(i) // i will be printed only when i is an even number
		}
	}
}
