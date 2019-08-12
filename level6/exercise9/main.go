package main

import "fmt"

func main() {
	f := func(xi []int) int {
		for _, v := range xi {
			fmt.Println(v)
		}
		return len(xi)
	}
	x := lengther(f, []int{
		1, 2, 3, 9,
	})
	fmt.Println(x)
}

func lengther(length func(xi []int) int, ii []int) int {
	return length(ii)
}
