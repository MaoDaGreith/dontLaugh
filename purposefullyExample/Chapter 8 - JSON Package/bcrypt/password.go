package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := "qwerty123"
	hashed, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(hashed)

	// 	loginPword1 := `qwerty123`

	// 	err = bcrypt.CompareHashAndPassword(hashed, []byte(loginPword1))
	// 	if err != nil {
	// 		fmt.Println("YOU CAN'T LOGIN")
	// 		return
	// 	}

	// 	fmt.Println("You're logged in")
	//
}
