package main

import (
	"encoding/json"
	"fmt"
)

type human struct {
	Gender      string
	Nationality string
	Mature      bool
}

func main() {
	h1 := human{
		Gender:      "Male",
		Nationality: "Tajik",
		Mature:      false,
	}

	h2 := human{
		Gender:      "Female",
		Nationality: "Armenian",
		Mature:      true,
	}

	s := []human{h1, h2}

	js, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(js))
}
