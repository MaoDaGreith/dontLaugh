package main

import "fmt"

type person struct {
	first_name  string
	last_name   string
	favIceCream []string
}

func main() {
	p1 := person{
		first_name:  "John",
		last_name:   "Connor",
		favIceCream: []string{"Vanilla", "Chai", "Cocoa", "Milkyway"},
	}

	p2 := person{
		first_name:  "Kevin",
		last_name:   "Bruster",
		favIceCream: []string{"Chocolate", "Banana", "Mango", "Strawberry"},
	}

	fmt.Println(p1.first_name, p1.last_name)
	for i, v := range p1.favIceCream {
		fmt.Println(i, v)
	}

	fmt.Println(p2.first_name, p2.last_name)
	for i, v := range p2.favIceCream {
		fmt.Println(i, v)
	}
}
