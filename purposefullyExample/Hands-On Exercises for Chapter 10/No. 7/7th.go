package main

import (
	"fmt"
	"sync"
)

// write a program that
// launches 10 goroutines
// each goroutine adds 10 numbers to a channel
// pull the numbers off the channel and print them

func main() {
	c := make(chan int)

	fanOut(c)
	print(c)

}

func print(c chan int) {
	for v := range c {
		fmt.Println(v)
	}
}

func fanOut(c chan int) {
	var wg sync.WaitGroup

	go func() {
		for i := 1; i <= 10; i++ {
			d := i
			wg.Add(1)
			go func() {
				for j := 1; j <= 10; j++ {
					c <- j * d
				}
				wg.Done()
			}()
			wg.Wait()
		}
		close(c)
	}()

}
