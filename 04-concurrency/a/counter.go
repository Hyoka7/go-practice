package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	a := 0
	for i := 0; i < 1000; i++ {
		wg.Go(func() {
			mutex.Lock()
			a += 1
			mutex.Unlock()
		})
	}
	wg.Wait()
	fmt.Println(a)
}
