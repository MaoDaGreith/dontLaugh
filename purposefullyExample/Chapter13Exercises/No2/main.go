package main

import (
	"fmt"
	"hash/example/Chapter13Exercises/No2/quote"
	"hash/example/Chapter13Exercises/No2/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
