package main

import "fmt"

func main() {
	func(s int) {
		fmt.Println("something", s)
	}(12234)
}
