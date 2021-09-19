package main

import "fmt"

type person struct {
	first_name string
	last_name  string
	age        int32
}

func main() {
	p1 := person{
		first_name: "Shrek",
		last_name:  "the Ogre",
		age:        27,
	}

	fmt.Println(p1)
	fmt.Println(p1.first_name, p1.last_name, p1.age)

}
