package main

import "fmt"

func main() {
	a := 4
	fmt.Printf("%d\t %b %#x\n", a, a, a)
	b := 4 << 1
	fmt.Printf("%d\t %b %#x\n", b, b, b)
	c := 4 >> 2
	fmt.Printf("%d\t %b %#x\n", c, c, c)
}
