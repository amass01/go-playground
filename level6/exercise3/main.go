package main

import "fmt"

func main() {
	deferd()
}

func deferd() {
	for i := 1; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("First this")
}
