package main

import (
	"fmt"
	"math"
)

type square struct {
	l int
	w int
}

type circle struct {
	r int
}

type shape interface {
	area() int
}

func (s square) area() int {
	return s.l * s.w
}

func (c circle) area() int {
	return int(3.14 * math.Sqrt(float64(c.r)))
}

func info(sh shape) {
	fmt.Println(sh.area())
}

func main() {
	s := square{
		l: 50,
		w: 50,
	}

	c := circle{
		r: 20,
	}

	fmt.Println(s.area())
	fmt.Println(c.area())
}
