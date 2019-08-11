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

	p2 := person{
		"Amirs",
		"Massarwas",
		[]string{"blue", "red"},
	}

	fmt.Println(p2)

	for _, v := range p2.favIce {
		fmt.Println(v)
	}
}
