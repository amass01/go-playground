package main

import "fmt"

func main() {
	xx := [][]string{{"James", "Bond", "Shaken, not stirred"}, {"Miss", "Moneypenny", "Helloooooo, James."}}
	fmt.Println(xx)

	for i, x := range xx {
		for j, v := range x {
			fmt.Println(i, j, v)
		}
	}
}
