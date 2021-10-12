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
	s := `[{"Gender":"Male","Nationality":"Tajik","Mature":false},{"Gender":"Female","Nationality":"Armenian","Mature":true}]`
	b := []byte(s)

	var humans []human
	err := json.Unmarshal(b, &humans)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(humans)
}
