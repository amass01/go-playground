package main

import "fmt"

func main() {
	a := [5]int{2, 0, 0, 9, 1}
	for i, v := range a {
		fmt.Printf("%v\t%v\n", i, v)
	}

	fmt.Printf("%T\n", a)
}
