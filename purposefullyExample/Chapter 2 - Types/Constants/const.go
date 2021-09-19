package main

import "fmt"

const a = 42      //untyped const
const b = 42.332  //untyped const
const c = "Shrek" //untyped const

/* the following version is also possible
const (
		a int = 42 			//typed const
		b float64 = 42.332 	//typed const
		c string = "Shrek" 	//typed const
)
*/

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}

// Note: Constants can be of 2 types typed and untyped

// regarding const float64 to int https://stackoverflow.com/a/58407938

// for more info: https://docs.google.com/document/d/1ckYpi6hcRkaBUEk975f54oGsHYHu7GhzOk7-nOrkNxo/edit#heading=h.f9fi1nfd9a3l
