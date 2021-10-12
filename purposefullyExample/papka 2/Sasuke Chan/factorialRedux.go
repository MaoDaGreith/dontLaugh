package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	// Seed the random number generator.
	rand.Seed(time.Now().UnixNano())
}

func main() {
	pop := make(<-chan int)

	pop = print(factorial())

	for v := range pop {
		fmt.Println(v)
	}
}

func factorial() chan int {
	var wg sync.WaitGroup
	d := make(chan int)
	go func() {
		for j := 1; j <= 100; j++ {
			b := 1
			wg.Add(1)
			go func() {
				for i := rand.Intn(10) + 1; i >= 1; i-- {
					b *= i
				}
				d <- b
				wg.Done()
			}()
			wg.Wait()
		}
		close(d)
	}()
	return d
}

func print(c <-chan int) <-chan int {
	d := make(chan int)
	go func() {
		for v := range c {
			d <- v + 2
		}
		close(d)
	}()

	return d
}

// func print(c <-chan int) {
// 	for v := range c {
// 		fmt.Println(v)
// 	}
// }
