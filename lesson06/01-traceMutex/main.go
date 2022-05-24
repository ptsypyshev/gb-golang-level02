package main

import (
	"fmt"
	"github.com/pkg/profile"
	"sync"
)

const (
	lenAlphabet        = 26
	shiftToFirstLetter = 65
)

type safeMap struct {
	m   map[int]string
	rwm *sync.RWMutex
}

func NewSafeMap() *safeMap {
	return &safeMap{
		m:   make(map[int]string),
		rwm: &sync.RWMutex{},
	}
}

func (sm safeMap) SetVal(val string) {
	sm.rwm.Lock()
	sm.m[len(sm.m)] = val
	sm.rwm.Unlock()
}

func (sm safeMap) GetVal(key int) (string, bool) {
	sm.rwm.RLock()
	val, ok := sm.m[key]
	sm.rwm.RUnlock()
	return val, ok
}

func (sm safeMap) PrintMap() {
	fmt.Println(sm.m)
}

func main() {
	defer profile.Start(profile.TraceProfile, profile.ProfilePath(".")).Stop()
	wg := sync.WaitGroup{}
	myMap := NewSafeMap()
	wg.Add(lenAlphabet)
	for i := 0; i < lenAlphabet; i++ {
		go func(i int) {
			defer wg.Done()
			myMap.SetVal(string(rune(i + shiftToFirstLetter)))
		}(i)
	}
	wg.Wait()
	myMap.PrintMap()
}
