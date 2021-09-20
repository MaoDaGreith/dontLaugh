package main

import "fmt"

type child struct {
	first string
	last  string
}

type human struct {
	child
	sex string
}

func (h human) voice() {
	fmt.Println("This method can see the human from:", h.first, h.last)
}

func main() {
	h1 := human{
		child: child{
			first: "Shrek",
			last:  "the Ogre",
		},
		sex: "little Man",
	}

	h1.voice()
}
