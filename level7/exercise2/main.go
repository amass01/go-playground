package main

import "fmt"

type person struct {
	first string
	last  string
}

func main() {
	p := person{
		first: "amir",
		last:  "mass",
	}
	fmt.Println(p)
	changeMe(&p)
	fmt.Println(p)
}

func changeMe(p *person) {
	p.first = "Asil"
}
