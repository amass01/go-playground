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

	// explicit use of struct fields names when creating an instance
	p2 := person{
		firstName: "Amirs",
		lastName:  "Massarwas",
		favIce:    []string{"blue", "red"},
	}

	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	for _, v := range m {
		fmt.Println(v)
	}

}
