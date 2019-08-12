package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p := person{
		first: `amir`,
		last:  `mass`,
		age:   27,
	}

	p.speak()
}

func (p person) speak() {
	fmt.Printf("My name is %v %v, and %v is my age!\n", p.first, p.last, p.age)
}
