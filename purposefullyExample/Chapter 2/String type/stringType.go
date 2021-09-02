package main

import "fmt"

func main() {
	s := "Shrek is the best movie" //assigning value to this variable
	fmt.Println(s)                 // displaying the string
	fmt.Printf("%T\n", s)          //displaying the Type of the variable
	b := []byte(s)                 // every string is a slice of bytes, that's why we able to convert the string into a slice
	fmt.Println(b)
	fmt.Printf("%T\n", b) // displaying the slice in decimals

	// looping through the slice representing the letter with the equivalent in UTF-8
	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i]) // %#U in order to show the equivalent of the letter in UTF-8
	}

	// displaying the index positions and the values of the s variable in hex
	fmt.Println()
	for i, v := range s {
		fmt.Printf("at index position %d we have hex %#x\n", i, v)
	}
}

// Note: strings are immutable and a sequence/slice of bytes

// For more info: https://docs.google.com/document/d/1ckYpi6hcRkaBUEk975f54oGsHYHu7GhzOk7-nOrkNxo/edit#heading=h.4s24qveq7prg
