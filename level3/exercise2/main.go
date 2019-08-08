package main

import "fmt"

func main() {
	for i := 65; i <= 90; i++ {
		j := 0
		for j < 3 {
			fmt.Printf("%#U\n", i)
			j++
		}
	}
}
