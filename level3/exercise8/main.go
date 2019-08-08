package main

import "fmt"

func main() {
	i := 2
	switch {
	case 2 == 2:
		fmt.Printf("%v\n", i)
	case 1 < 3:
		fmt.Println("Shouldn't print!")
	default:
		fmt.Println("Shouldn't print as well")
	}
}
