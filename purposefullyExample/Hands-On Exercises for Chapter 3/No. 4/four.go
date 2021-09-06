package main

import "fmt"

func main() {
	year := 1995
	for {
		fmt.Println(year)
		year++
		if year > 2021 {
			break
		}
	}
}
