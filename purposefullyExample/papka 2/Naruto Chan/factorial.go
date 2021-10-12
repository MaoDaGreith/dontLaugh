package main

import "fmt"

func main() {
	pop := make(chan int)
	rec := make(chan int)

	go func() {
		pop <- 7
		close(pop)
	}()
	rec = factorial(pop)
	fmt.Println(<-rec)

	//close(rec)
}

func factorial(c <-chan int) chan int {
	d := make(chan int)
	b := 1
	go func() {
		for i := <-c; i >= 1; i-- {
			b *= i
		}
		d <- b
		close(d)
	}()
	return d
}
