package main

import "fmt"

func main() {
	daD(func() string {
		return "lel"
	})
}

func daD(f func() string) {
	fmt.Println(f())
}
