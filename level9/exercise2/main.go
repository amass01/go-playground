package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("Hi, I'm", p.first, p.last, ", and", p.age, "is my age")
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p := person{
		first: "Amir",
		last:  "Mass",
		age:   27,
	}

	// saySomething(p) - won't work
	saySomething(&p)
}
