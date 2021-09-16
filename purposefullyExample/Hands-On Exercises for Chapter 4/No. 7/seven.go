package main

import "fmt"

func main() {
	s := [][]string{{"James", "Bond", "Shaken, not stirred"}, {"Miss", "Moneypenny", "Helloooooo, James."}}

	for i, v := range s {
		fmt.Println("record: ", i)
		for j, val := range v {
			fmt.Printf("\t\tindex position: %v \t value: %v\n", j, val)
		}
	}

}
