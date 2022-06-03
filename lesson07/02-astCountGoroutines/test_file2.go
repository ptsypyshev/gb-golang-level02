package main

import (
	"fmt"
	"sync"
)

func another() {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 5; i++ {
		go func(v int, wg *sync.WaitGroup) {
			fmt.Println(v)
			wg.Done()
		}(i, &wg)
		go func(v int, wg *sync.WaitGroup) {
			fmt.Println(v)
			wg.Done()
		}(i, &wg)
	}
	wg.Wait()
}
