package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go printAnotherSomething()
	go printSomething()
	wg.Wait()
}

func printSomething() {
	fmt.Println("This is something")
	wg.Done()
}

func printAnotherSomething() {
	fmt.Println("Another something printed")
	wg.Done()
}
