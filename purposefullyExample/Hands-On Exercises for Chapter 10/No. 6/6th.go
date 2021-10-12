package main

import "fmt"

// write a program that
// puts 100 numbers to a channel
// pull the numbers off the channel and print them

func main() {
	c := send()

	receive(c)
}

func send() <-chan int {
	c := make(chan int)
	go func() {
		for i := 1; i <= 100; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
