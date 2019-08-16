package main

import (
	"fmt"
)

// second gou solution
// func main() {
// 	c := make(chan int)

// 	go func() {
// 		c <- 42
// 	}()

// 	fmt.Println(<-c)
// }

func main() {
	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)
}
