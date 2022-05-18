package main

import (
	"fmt"
	"time"
)

func main() {
	var workers = make(chan struct{}, 100)
	var counter int
	for i := 1; i <= 1000; i++ {
		workers <- struct{}{}
		go func(job int) {
			defer func() {
				<-workers // Read from channel to unlock another waiting goroutines
			}()
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("Job number: %d; Counter: %d, Number of run goroutines %d\n", job, counter, len(workers))
			counter++
		}(i)
	}
	time.Sleep(2 * time.Second)
	fmt.Printf("I've counted %d goroutines.\n", counter)
}
