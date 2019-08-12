package main

import "fmt"

func main() {
	f := idGenerator()

	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

func idGenerator() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
