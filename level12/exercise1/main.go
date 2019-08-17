package main

import (
	"fmt"

	"github.com/amassarwi/go-playground/level12/dog"
)

func main() {
	myAge := 27.5

	fmt.Printf("And my dog age is: %v\n", dog.Years(myAge))
}
