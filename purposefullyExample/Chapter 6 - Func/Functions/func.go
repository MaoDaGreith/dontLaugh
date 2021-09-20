package main

import "fmt"

func main() {
	bar()
	foo("Shrek")
	s := woo("Peanut butter on my...")
	fmt.Println(s)
}

func bar() {
	fmt.Println("This func has no parameters and no returns")
}

func foo(s string) {
	fmt.Println("This function has a parameter:", s)
}

func woo(s string) (r string) {
	return fmt.Sprint("This is a return with: ", s)
}
