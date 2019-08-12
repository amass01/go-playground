package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("I was passed to another function")
	}
	execute(f)
}

func execute(toRun func()) {
	toRun()
}
