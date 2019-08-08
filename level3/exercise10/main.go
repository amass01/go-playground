package main

import "fmt"

func main() {
	// fmt.Println( true && true ) - should print true
	fmt.Println(true && false)
	// fmt.Println(true || true)  - should print true
	fmt.Println(true || false)
	fmt.Println(!true)
}
