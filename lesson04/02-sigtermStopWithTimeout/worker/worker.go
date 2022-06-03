// Package worker can be used as a worker in pool
package worker

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

const doPeriod = 100

// Worker is an empty struct which can be used as worker unit
type Worker struct {
}

// NewWorker is initialization method of new worker
func NewWorker() *Worker {
	return &Worker{}
}

// Run method can be used to run another func with parent context
func (w *Worker) Run(ctx context.Context) {
	ticker := time.NewTicker(doPeriod * time.Millisecond)

	go func() {
		for {
			select {
			case <-ctx.Done():
				ticker.Stop()
				return
			case <-ticker.C:
				resp := Pinger()
				fmt.Println(resp)
			}
		}
	}()
}

// Pinger simple mock function emulates ping request/response
func Pinger() string {
	randTimeInMilliSecond := time.Duration(rand.Intn(100))
	time.Sleep(randTimeInMilliSecond * time.Millisecond)
	return fmt.Sprintf("%s ping is sucessful, response time is %d ms.", time.Now().Format("2006/01/02 15:04:05"), randTimeInMilliSecond+doPeriod)
}
