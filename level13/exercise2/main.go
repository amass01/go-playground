package main

import (
	"fmt"

	"github.com/amassarwi/go-playground/level13/exercise2/quote"
	"github.com/amassarwi/go-playground/level13/exercise2/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
