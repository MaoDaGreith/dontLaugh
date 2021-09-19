package main

import "fmt"

type person struct {
	first_name string
	last_name  string
	age        int
}

type human struct {
	person
	sex         string
	nationality string
}

func main() {
	h1 := human{
		person: person{
			first_name: "Donald",
			last_name:  "the Duck",
			age:        25,
		},
		sex:         "male",
		nationality: "american",
	}

	fmt.Println(h1)
	fmt.Println(h1.first_name, h1.last_name, h1.age, h1.nationality, h1.sex)

}
