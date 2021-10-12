package main

import (
	"fmt"
	"runtime"
)

func fibonacci(c, quit chan int) {
	fmt.Println(runtime.NumGoroutine())
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	fmt.Println("before go func")
	go func() {
		fmt.Println("before")
		for i := 0; i < 10; i++ {
			fmt.Println("before sending")
			fmt.Println("Trying to send", <-c)
		}
		quit <- 1
	}()
	fmt.Println("first goroutine shown", runtime.NumGoroutine())
	fmt.Println("Sending for the first time")
	fibonacci(c, quit)
}
