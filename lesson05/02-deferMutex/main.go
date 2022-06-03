package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	var mu = sync.Mutex{}
	var wg = sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mu.Unlock()
				wg.Done()
			}()
			mu.Lock()
			counter++
		}()
	}

	wg.Wait()
	fmt.Printf("I've counted %d goroutines.\n", counter)
}
