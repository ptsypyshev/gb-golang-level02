package main

import "sync"

type floatMap struct {
	set map[float64]struct{}
	m   *sync.Mutex
	rwm *sync.RWMutex
}

func NewFloatMap() *floatMap {
	return &floatMap{
		set: map[float64]struct{}{},
		m:   &sync.Mutex{},
		rwm: &sync.RWMutex{},
	}
}

func (f floatMap) Get(k float64) {
	f.m.Lock()
	_ = f.set[k]
	f.m.Unlock()
}

func (f floatMap) GetRW(k float64) {
	f.rwm.RLock()
	_ = f.set[k]
	f.rwm.RUnlock()
}

func (f floatMap) Set(k float64) {
	f.m.Lock()
	f.set[k] = struct{}{}
	f.m.Unlock()
}

func (f floatMap) SetRW(k float64) {
	f.rwm.Lock()
	f.set[k] = struct{}{}
	f.rwm.Unlock()
}

func main() {
	fMap := NewFloatMap()
	fMap.Set(1.2)
	fMap.Get(1.4)
	fMap.SetRW(1.2)
	fMap.GetRW(1.4)
}
