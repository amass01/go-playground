package main

import "fmt"

func main() {
	fmt.Println(foo())
	fmt.Println(bar())
}

func foo() int {
	return 27
}

func bar() (string, int) {
	return "I mean My Age", 27
}
