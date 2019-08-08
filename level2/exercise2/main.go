package main

import "fmt"

func main() {
	a := 5 == 3
	b := 3 <= 1.5
	c := a != b
	d := 0 < -1
	e := 1 > -42
	fmt.Println(a, b, c, d, e)
}
