package main

import (
	"errors"
	"fmt"
)

type customErr struct {
	err error
}

func (ce customErr) Error() string {
	return fmt.Sprintf("This is a custom err: bruh %v", ce.err)
}

func main() {
	c1 := customErr{err: errors.New("Dis is in struct")}
	foo(c1)
}

func foo(e error) {
	fmt.Println(e)
}
