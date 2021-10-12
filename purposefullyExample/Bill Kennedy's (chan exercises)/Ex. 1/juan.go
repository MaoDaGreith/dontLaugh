// Write a program where two goroutines pass an integer back and forth
// ten times. Display when each goroutine receives the integer. Increment
// the integer with each pass. Once the integer equals ten, terminate
// the program cleanly.
package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Add imports.

func main() {
	c := make(chan int)
	// Create an unbuffered channel.

	var wg sync.WaitGroup
	wg.Add(2)
	// Create the WaitGroup and add a count
	// of two, one for each goroutine.

	// Launch the goroutine and handle Done.
	go func() {
		fmt.Println("First goroutine received the integer")
		goroutine(c)
		wg.Done()
	}()

	// Launch the goroutine and handle Done.
	go func() {
		fmt.Println("Second goroutine received the integer")
		goroutine(c)
		wg.Done()
	}()

	// Send a value to start the counting.
	c <- 1

	// Wait for the program to finish.
	wg.Wait()
}

// goroutine simulates sharing a value.
func goroutine(c chan int) {
	fmt.Println(runtime.NumGoroutine())
	for {

		// Wait for the value to be sent.
		// If the channel was closed, return.
		v, ok := <-c
		if !ok {
			fmt.Println("Channel closed")
			return
		}

		// Display the value.
		fmt.Println(v)
		// Terminate when the value is 10.
		if v == 10 {
			fmt.Println("Closing Channel")
			close(c)
			return
		}

		// Increment the value and send it
		// over the channel.
		c <- v + 1
	}
}
