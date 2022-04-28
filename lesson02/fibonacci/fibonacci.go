// Package fibonacci can be used to calculate Fibonacci numbers
//It has 2 variants of realization - naive and optimized.
package fibonacci

var fibMap = map[int]int{
	0: 0,
	1: 1,
}

// FibWithoutCache returns value of Fibonacci number (argument). Naive version of func.
func FibWithoutCache(n int) int {
	if n < 2 {
		return n
	}
	return FibWithoutCache(n-2) + FibWithoutCache(n-1)
}

// FibWithCache returns value of Fibonacci number (argument). Optimized version of func.
func FibWithCache(n int) int {
	if val, ok := fibMap[n]; ok {
		return val
	}

	fibMap[n] = FibWithCache(n-2) + FibWithCache(n-1)

	return fibMap[n]
}
