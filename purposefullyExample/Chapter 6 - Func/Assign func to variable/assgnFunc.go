package main

import "fmt"

func main() {
	s := func(a int) (r string) {
		fmt.Printf("%v is the best Shrek EVER", a)
		if a > 0 {
			return "Syke"
		} else {
			return "dam"
		}
	}

	s(0)
}
