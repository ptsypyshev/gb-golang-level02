package main

import (
	"context"
	"fmt"
	"github.com/ptsypyshev/gb-golang-level02/lesson04/02-sigtermStopWithTimeout/worker"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// is a handler of signal
func sigHandler(cancel context.CancelFunc) {
	sigChannel := make(chan os.Signal)
	signal.Notify(sigChannel, syscall.SIGTERM, syscall.SIGINT)

	<-sigChannel
	fmt.Printf("%s Get SIGTERM, try to stop.\n", time.Now().Format("2006/01/02 15:04:05"))
	cancel()
	time.Sleep(time.Second)
	log.Fatal("Application is killed!\n")

}

func main() {
	// Create new worker object
	w := worker.NewWorker()

	// Create new context
	ctx, cancel := context.WithCancel(context.Background())

	// Run worker
	w.Run(ctx)

	// Start handling of syscalls
	sigHandler(cancel)
}
