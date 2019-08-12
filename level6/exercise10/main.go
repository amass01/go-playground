package main

import "fmt"

func main() {
	x := 5
	encloser(x)
}

func encloser(x int) {
	fmt.Println(x)

	{
		x := "test"
		fmt.Println(x)
	}
}
