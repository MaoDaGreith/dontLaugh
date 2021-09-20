package main

import "fmt"

func main() {
	v := returns()

	fmt.Println(v())
}

func returns() func() string {
	return func() string {
		return "string"
	}
}
