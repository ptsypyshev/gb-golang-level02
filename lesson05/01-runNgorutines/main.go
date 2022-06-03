package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

// worker emulates random execution time of function
func worker(job int, wg *sync.WaitGroup) {
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())

	fmt.Printf("Job %d is started.\n", job)

	randWorkingTime := time.Duration(rand.Intn(10)) * time.Second
	time.Sleep(randWorkingTime)

	fmt.Printf("Job %d is finished, сompleted in %.f seconds.\n", job, randWorkingTime.Seconds())
}

func main() {
	var (
		n  int
		wg = sync.WaitGroup{}
	)

	fmt.Print("How much goroutines should I start? ")
	if _, err := fmt.Scan(&n); err != nil {
		log.Fatalf("Cannot contunue: %v", err)
	}

	wg.Add(n)
	for i := 1; i <= n; i++ {
		go worker(i, &wg)
	}

	wg.Wait()
}
