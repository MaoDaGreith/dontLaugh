package main

import "fmt"

func main() {
	s := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr:`:          []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	for i, v := range s {
		fmt.Println("Index of: ", i)
		for j, val := range v {
			fmt.Printf("\t\t%v\t%v\n", j, val)
		}
	}
}
