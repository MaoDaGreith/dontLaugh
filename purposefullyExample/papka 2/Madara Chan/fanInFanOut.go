package main

import (
	"fmt"
	"sync"
)

func main() {

	in := gen()

	c01 := factorial(in)
	c02 := factorial(in)
	c03 := factorial(in)
	c04 := factorial(in)
	c05 := factorial(in)
	c06 := factorial(in)
	c07 := factorial(in)
	c08 := factorial(in)
	c09 := factorial(in)
	c00 := factorial(in)

	var y int
	for n := range merge(c00, c01, c02, c03, c04, c05, c06, c07, c08, c09) {
		y++
		fmt.Println(y, "\t", n)
	}
}

func gen() <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)
	go func() {
		for i := 0; i < 1000; i++ {
			wg.Add(1)
			go func() {
				for j := 3; j < 13; j++ {
					out <- j
				}
				wg.Done()
			}()
			wg.Wait()
		}
		close(out)
	}()
	return out
}

func factorial(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- fact(n)
		}
		close(out)
	}()
	return out
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

func merge(c ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(c))
	for _, v := range c {
		go output(v)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

/*
CHALLENGE #1:
-- Change the code above to execute 1,000 factorial computations concurrently and in parallel.
-- Use the "fan out / fan in" pattern
CHALLENGE #2:
WATCH MY SOLUTION BEFORE DOING THIS CHALLENGE #2
-- While running the factorial computations, try to find how much of your resources are being used.
-- Post the percentage of your resources being used to this discussion: https://goo.gl/BxKnOL
*/
