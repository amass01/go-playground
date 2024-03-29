package main

import "fmt"

func main() {
	c := make(chan int)
	for i := 1; i <= 10; i++ {
		go func() {
			for j := 1; j <= 10; j++ {
				c <- j
			}
		}()
	}

	for k := 0; k < 100; k++ {
		fmt.Println(k, <-c)
	}
}
