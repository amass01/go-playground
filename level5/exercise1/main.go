package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	favIce    []string
}

func main() {
	p1 := person{
		"Amir",
		"Massarwa",
		[]string{"pink", "green"},
	}

	fmt.Println(p1)

	// explicit use of struct fields names when creating an instance
	p2 := person{
		firstName: "Amirs",
		lastName:  "Massarwas",
		favIce:    []string{"blue", "red"},
	}

	fmt.Println(p2)

	for _, v := range p2.favIce {
		fmt.Println(v)
	}
}
