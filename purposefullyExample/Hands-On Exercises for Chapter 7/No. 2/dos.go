package main

import "fmt"

type person struct {
	first string
	last  string
}

func main() {
	p1 := person{
		first: "Kim",
		last:  "Possible",
	}
	changeMe(&p1)
	fmt.Println(p1)
}

func changeMe(p *person) {
	p.first = "Shrek"
	p.last = "the Ogre"
	// The other way to dereference a value from a struct is the following example
	//(*p).first = "Shrek"
	//(*p).last = "the Ogre"
}
