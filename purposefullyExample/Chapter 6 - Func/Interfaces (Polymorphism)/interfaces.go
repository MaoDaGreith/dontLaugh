package main

import "fmt"

type episode struct {
	title  string
	number int
}

type anime struct {
	episode
	season int
}

func (a anime) voice() {
	fmt.Println("This method can see the human from:", a.number, a.season)
}

func (e episode) voice() {
	fmt.Println("This method can see the human from:", e.title, e.number)
}

type content interface {
	voice()
}

func separate(c content) {
	switch c.(type) {
	case anime:
		fmt.Println("This is ANIME TITTIES")
	case episode:
		fmt.Println("This is episodezzzz")
	}
}

func main() {
	a1 := anime{
		episode: episode{
			title:  "Shrek",
			number: 5,
		},
		season: 3,
	}

	e1 := episode{
		title:  "Thomas the engine",
		number: 65,
	}

	a1.voice()
	e1.voice()

	separate(a1)
	separate(e1)

}
