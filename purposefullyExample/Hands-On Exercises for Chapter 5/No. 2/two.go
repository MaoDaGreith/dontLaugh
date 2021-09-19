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

	m := map[string]person{
		p1.last_name: p1,
		p2.last_name: p2,
	}

	for _, v := range m {
		fmt.Println(v.first_name)
		fmt.Println(v.last_name)
		for i, val := range v.favIceCream {
			fmt.Println(i, val)
		}
		fmt.Println("-------------")
	}
}
