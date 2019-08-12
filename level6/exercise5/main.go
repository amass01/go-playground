package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (s square) area() float64 {
	return s.length * s.length
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(sh shape) {
	fmt.Println(sh.area())
}

func main() {
	s := square{
		length: 50,
	}

	c := circle{
		radius: 20,
	}

	fmt.Println(s.area())
	fmt.Println(c.area())
}
