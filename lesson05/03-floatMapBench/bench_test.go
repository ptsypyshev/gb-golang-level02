package main

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

const numGoroutines = 10

func mockRandData() []float64 {
	rand.Seed(time.Now().UnixNano())
	var result = make([]float64, 0, numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		result = append(result, rand.Float64())
	}
	//fmt.Println(result, len(result))
	return result
}

func BenchmarkMutex10w_90r(b *testing.B) {
	wg := &sync.WaitGroup{}
	benchSet := NewFloatMap()
	randSlice := mockRandData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < numGoroutines; j++ {
			wg.Add(1)
			if j%10 != 0 {
				go func(j int) {
					benchSet.Get(randSlice[j])
					wg.Done()
				}(j)
			} else {
				go func(j int) {
					benchSet.Set(randSlice[j])
					wg.Done()
				}(j)
			}
		}
		wg.Wait()
	}
}

func BenchmarkRWMutex10w_90r(b *testing.B) {
	wg := &sync.WaitGroup{}
	benchSet := NewFloatMap()
	randSlice := mockRandData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < numGoroutines; j++ {
			wg.Add(1)
			if j%10 != 0 {
				go func(j int) {
					benchSet.GetRW(randSlice[j])
					wg.Done()
				}(j)
			} else {
				go func(j int) {
					benchSet.SetRW(randSlice[j])
					wg.Done()
				}(j)
			}
		}
		wg.Wait()
	}
}

func BenchmarkMutex50w_50r(b *testing.B) {
	wg := &sync.WaitGroup{}
	benchSet := NewFloatMap()
	randSlice := mockRandData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < numGoroutines; j++ {
			wg.Add(1)
			if j%2 == 0 {
				go func(j int) {
					benchSet.Get(randSlice[j])
					wg.Done()
				}(j)
			} else {
				go func(j int) {
					benchSet.Set(randSlice[j])
					wg.Done()
				}(j)
			}
		}
		wg.Wait()
	}
}

func BenchmarkRWMutex50w_50r(b *testing.B) {
	wg := &sync.WaitGroup{}
	benchSet := NewFloatMap()
	randSlice := mockRandData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < numGoroutines; j++ {
			wg.Add(1)
			if j%2 == 0 {
				go func(j int) {
					benchSet.GetRW(randSlice[j])
					wg.Done()
				}(j)
			} else {
				go func(j int) {
					benchSet.SetRW(randSlice[j])
					wg.Done()
				}(j)
			}
		}
		wg.Wait()
	}
}

func BenchmarkMutex90w_10r(b *testing.B) {
	wg := &sync.WaitGroup{}
	benchSet := NewFloatMap()
	randSlice := mockRandData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < numGoroutines; j++ {
			wg.Add(1)
			if j%10 == 0 {
				go func(j int) {
					benchSet.Get(randSlice[j])
					wg.Done()
				}(j)
			} else {
				go func(j int) {
					benchSet.Set(randSlice[j])
					wg.Done()
				}(j)
			}
		}
		wg.Wait()
	}
}

func BenchmarkRWMutex90w_10r(b *testing.B) {
	wg := &sync.WaitGroup{}
	benchSet := NewFloatMap()
	randSlice := mockRandData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < numGoroutines; j++ {
			wg.Add(1)
			if j%10 == 0 {
				go func(j int) {
					benchSet.GetRW(randSlice[j])
					wg.Done()
				}(j)
			} else {
				go func(j int) {
					benchSet.SetRW(randSlice[j])
					wg.Done()
				}(j)
			}
		}
		wg.Wait()
	}
}
