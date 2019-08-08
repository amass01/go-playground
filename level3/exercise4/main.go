package main

import "fmt"

func main() {
	year := 1991
	for {
		if year > 2019 {
			break
		}
		fmt.Printf("%v\n", year)
		year++
	}
}
