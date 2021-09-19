package main

import "fmt"

func main() {
	h1 := struct {
		hp     int
		armor  int
		weapon string
		class  string
	}{
		hp:     32,
		armor:  76,
		weapon: "Nail Gun",
		class:  "Soldier",
	}

	fmt.Println(h1)
}
