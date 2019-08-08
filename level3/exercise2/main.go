package main

import "fmt"

func main() {
	for i := 65; i <= 90; i++ {
		fmt.Printf("%v\n", i)
		j := 0
		for j < 3 {
			fmt.Printf("\t%#U\n", i)
			j++
		}
	}
}
