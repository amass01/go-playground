package main

import "fmt"

func main() {
	anonymousPerson := struct {
		name string
		age  int
	}{
		"Amirmass",
		23,
	}

	fmt.Println(anonymousPerson)
}
