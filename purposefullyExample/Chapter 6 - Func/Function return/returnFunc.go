package main

import "fmt"

func main() {
	fmt.Println(ret()())
}

func ret() func() string {
	return func() string {
		return "Everytime I look at my di... DIO!!!"
	}
}
