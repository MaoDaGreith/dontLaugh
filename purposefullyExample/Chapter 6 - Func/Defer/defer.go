package main

import "fmt"

func main() {
	defer sheesh()
	done()

}

func done() {
	fmt.Println("Syke that's the wrong Number.")
}

func sheesh() {
	fmt.Println("Is this the real life?")
}
