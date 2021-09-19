package main

import "fmt"

func main() {
	p1 := struct {
		first_name  string
		last_name   string
		favIceCream []string
		friends     map[string]int
	}{
		first_name:  "John",
		last_name:   "Connor",
		favIceCream: []string{"Vanilla", "Chai", "Cocoa", "Milkyway"},
		friends: map[string]int{
			"Moneypenny": 555,
			"Q":          777,
			"M":          888,
		},
	}

	fmt.Println(p1.first_name)
	fmt.Println(p1.last_name)
	fmt.Println(p1.favIceCream)

	for k, v := range p1.favIceCream {
		fmt.Println(k, v)
	}

	for i, val := range p1.friends {
		fmt.Println(i, val)
	}
}
