// Write a program that uses a fan out pattern to generate 100 random numbers
// concurrently. Have each goroutine generate a single random number and return
// that number to the main goroutine over a buffered channel. Set the size of
// the buffered channel so no send ever blocks. Don't allocate more capacity
// than you need. Have the main goroutine store each random number it receives
// in a slice. Print the slice values then terminate the program.
package main

// Add imports.
import (
	"fmt"
	"math/rand"
	"time"
)

// Declare constant for number of goroutines.
const routines = 100

func init() {
	// Seed the random number generator.
	rand.Seed(time.Now().UnixNano())
}

func main() {
	c := make(chan int, routines)
	// Create the buffered channel with room for
	// each goroutine to be created.

	// Iterate and launch each goroutine.
	// Create an anonymous function for each goroutine that
	// generates a random number and sends it on the channel.
	for i := 1; i <= routines; i++ {
		go func() {
			c <- rand.Intn(500)
		}()
	}

	// Create a variable to be used to track received messages.
	// Set the value to the number of goroutines created.
	d := routines

	// Iterate receiving each value until they are all received.
	// Store them in a slice of ints.
	var num []int
	for d > 0 {
		num = append(num, <-c)
		d--
	}

	fmt.Println(num)
	// Print the values in our slice.
}
