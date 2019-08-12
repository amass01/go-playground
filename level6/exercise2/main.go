package main

import "fmt"

func main() {
	fmt.Println("The sum of 1, 4, 6, 3, 6, 2, 3: is", foo(1, 4, 6, 3, 6, 2, 3))
}

func foo(nums ...int) int {
	s := 0
	for _, v := range nums {
		s += v
	}
	return s
}
