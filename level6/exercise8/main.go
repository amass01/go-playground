package main

import "fmt"

var nextID func()

func main() {
	nextID := idGenerator()

	fmt.Println(nextID())
	fmt.Println(nextID())
	fmt.Println(nextID())
	fmt.Println(nextID())
}

func idGenerator() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
