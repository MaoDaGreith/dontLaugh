package main

import (
	"fmt"
)

func main() {
	// solution 2
	d := make(chan int, 2)

	d <- 3

	c := make(chan int)
	// solution 1
	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
	fmt.Println(<-d)
}

// Note: for buffered channel there is no need to launch a go routine fi there is still space in it
// Goroutine launched is needed only for unbuffered channels
