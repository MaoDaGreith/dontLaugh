package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)
	d := make(chan int)
	fanin := make(chan int)

	// send
	go populate(c, d)

	// receive
	go fanIn(c, d, fanin)

	for v := range fanin {
		fmt.Println(v)
	}

}

func populate(c, d chan<- int) {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			c <- i
		} else {
			d <- i
		}
	}
	close(c)
	close(d)
}

func fanIn(c, d <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range c {
			fanin <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range d {
			fanin <- v
		}
		wg.Done()
	}()
	wg.Wait()
	close(fanin)
}
