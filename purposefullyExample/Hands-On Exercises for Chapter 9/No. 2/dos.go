package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p *person) speak() {
	fmt.Println(p.first, p.last, p.age)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		first: "Shrek",
		last:  "the Ogre",
		age:   27,
	}

	saySomething(&p1)

	//saySomething(p1)

	p1.speak()
}
