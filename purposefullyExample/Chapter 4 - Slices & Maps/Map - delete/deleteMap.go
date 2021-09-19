package main

import "fmt"

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)

	m["Shrek"] = 32

	delete(m, "James")
	fmt.Println(m)

	delete(m, "Ian Fleming")
	fmt.Println(m)
	fmt.Println(m["Ian Fleming"])

	fmt.Println(m["Miss Moneypenny"])
	fmt.Println(m["Shrek"])

	if v, ok := m["Miss Moneypenny"]; ok {
		fmt.Println("value:", v)
		delete(m, "Miss Moneypenny")
	}

	fmt.Println(m)

}
