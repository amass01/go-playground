package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	counter := 0
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			var m sync.Mutex
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
