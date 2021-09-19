package main

import "fmt"

func main() {
	// % <- modulus operator (used to divide 2 numbers and get the remainder of the two if there is any)

	// example of using "break"
	for i := 0; i < 10; i++ {
		if i == 7 { // when i is equal to 7 to statement between the brackets will be executed
			break // break statement is executed because i is equal to 7
			// break statement exits the loop if the condition is met
		}
		fmt.Println(i)
	}
	fmt.Printf("end of loop 1\n\n")

	// example of using "continue"
	for i := 0; i < 20; i++ {
		if i%2 == 0 { // we are using modulus to get the remainder in order to display only the odd numbers
			continue // continue is executed due the the previous condition being met
			// continue statement skips over whatever there is after the statement going back to the beginning of the loop (i gets incremented by 1 and line 24 is not executed)
		}
		fmt.Println(i)
	}
	fmt.Printf("end of loop 2\n\n")
	/*			Other forms of for

		// while-ish
		for i <= 3 {
	        fmt.Println(i)
	        i = i + 1
	    }

		// do-while-ish
		code from 9-15

		// just a for
		x := 1
		for {
			x++
			if x > 100 {
				break
			}

			if x%2 != 0 {
				continue
			}

		fmt.Println(x)

	}
	*/
}
