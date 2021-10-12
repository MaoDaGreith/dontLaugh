package main

import (
	"fmt"
)

var count int

func main() {
	inChan := make(chan int)
	inChan2 := make(chan int)
	go incrementor(inChan)
	print(inChan)
	go incrementor(inChan2)
	print(inChan2)

	fmt.Println("Final Counter:", count)
	panic("stacks?")
}

func incrementor(c chan int) {
	count = count + 1
	for i := 1; i <= 20; i++ {
		c <- i
	}
	close(c)
}

func print(c <-chan int) {
	for v := range c {
		fmt.Println("Process: ", count, "printing: ", v)
	}
}

func increments(n int) {

}

/*
CHALLENGE #1
-- Take the code from above and change it to use channels instead of wait groups and atomicity
*/
