package main

import (
	"fmt"
	"github.com/pkg/profile"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	defer profile.Start(profile.TraceProfile, profile.ProfilePath(".")).Stop()
	var (
		goroutinesCounter int
		m                 = sync.Mutex{}
	)
	for i := 0; i < 10; i++ {
		go func() {
			m.Lock()
			goroutinesCounter++
			m.Unlock()
			fmt.Printf("I'm goroutine %d!\n", goroutinesCounter)
			for {

			}
		}()
	}
	for i := 0; ; i++ {
		if i%1e5 == 0 {
			fmt.Println("I'm main goroutine!")
			runtime.Gosched()
		}
		if i%1e8 == 0 {
			time.Sleep(time.Millisecond)
			fmt.Println("Main is finished...")
			return
		}
	}
}
