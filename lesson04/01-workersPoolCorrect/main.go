package main

import (
	"fmt"
	"time"
)

// Function to start worker
func worker(pool chan struct{}, res chan int, job int) {
	defer func() {
		<-pool // Read from channel to unlock another waiting goroutines
	}()
	var result int
	for {
		select {
		case result = <-res:
			//fmt.Printf("Job number: %d; Counter: %d, Number of run goroutines %d\n", job, result, len(pool))
			result++
			res <- result
			return
		}
	}
}

func main() {
	// Pull of workers, max parallel workers is 100
	workers := make(chan struct{}, 100)
	defer close(workers)

	// Sync channel for integers
	intCh := make(chan int)
	defer close(intCh)

	//Init counter value
	go func() {
		intCh <- 0
	}()

	// Run 1000 workers
	for i := 1; i <= 1000; i++ {
		workers <- struct{}{}
		go worker(workers, intCh, i)
	}

	time.Sleep(time.Second)
	fmt.Printf("I've counted %d goroutines.\n", <-intCh)
}
