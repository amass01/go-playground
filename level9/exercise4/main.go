package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var m sync.Mutex
	counter := 0
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			m.Lock()
			v := counter
			v++
			counter = v
			m.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
