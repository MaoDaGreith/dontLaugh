package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)
	d := make(chan int)
	go populate(c)

	go fanOut(c, d)

	for v := range d {
		fmt.Println(v)
	}
}

func populate(c chan<- int) {
	for i := 1; i <= 100; i++ {
		c <- i * i
	}
	close(c)
}

func fanOut(c, d chan int) {
	var wg sync.WaitGroup
	for v := range c {
		wg.Add(1)
		go func(v2 int) {
			d <- v2
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(d)
}
