package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	wg.Add(1)
	go foo()
	wg.Add(1)
	go bar()
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Wait()
}

func foo() {
	for i := 0; i < 99; i++ {
		fmt.Println(i * i)
	}
	fmt.Println("Done")
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Done()
}

func bar() {
	for i := 0; i < 99; i++ {
		fmt.Println("Something")
	}
	fmt.Println("Done")
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Done()
}
