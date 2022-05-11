package fibonacci

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var resultMap = map[int]int{
	0:  0,
	1:  1,
	2:  1,
	3:  2,
	4:  3,
	5:  5,
	6:  8,
	7:  13,
	8:  21,
	9:  34,
	10: 55,
	11: 89,
	12: 144,
	13: 233,
	14: 377,
	15: 610,
	16: 987,
	17: 1597,
	18: 2584,
	19: 4181,
	20: 6765,
	21: 10946,
	22: 17711,
	23: 28657,
	24: 46368,
	25: 75025,
	26: 121393,
	27: 196418,
	28: 317811,
	29: 514229,
	30: 832040,
}

func TestFibWithCache(t *testing.T) {
	for i := 0; i < len(resultMap); i++ {
		assert.Equal(t, resultMap[i], FibWithCache(i))
	}
}

func TestFibWithoutCache(t *testing.T) {
	for i := 0; i < len(resultMap); i++ {
		assert.Equal(t, resultMap[i], FibWithoutCache(i))
	}
}

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//FibWithCache(30)
		FibWithoutCache(30)
	}
}

func ExampleFibWithCache() {
	f := FibWithCache(10)
	fmt.Println(f)

	//Output: 55
}

func ExampleFibWithoutCache() {
	f := FibWithoutCache(10)
	fmt.Println(f)

	//Output: 55
}
