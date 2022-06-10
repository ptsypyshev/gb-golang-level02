package main

import (
	"fmt"
)

const (
	numGoroutines    = 1000
	numWorkersInPool = 100
)

func main() {
	var (
		workers = make(chan struct{}, numWorkersInPool)
		counter int
		//m       sync.Mutex
		//wg      sync.WaitGroup
	)
	//wg.Add(numGoroutines)
	for i := 1; i <= numGoroutines; i++ {
		workers <- struct{}{}
		go func(job int) {
			defer func() {
				<-workers // Read from channel to unlock another waiting goroutines
				//wg.Done()
			}()
			//m.Lock()
			counter++
			//m.Unlock()
		}(i)
	}
	//wg.Wait()
	fmt.Printf("I've counted %d goroutines.\n", counter)
}
