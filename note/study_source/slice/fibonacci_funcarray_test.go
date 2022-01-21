package slice

import (
	"testing"
)

func TestFibonacciFuncarray(t *testing.T) {
	term := 15
	result := fibarray(term)
	for ix, fib := range result {
		t.Logf("The %d-th Fibonacci number is %d\n", ix, fib)
	}
}
func fibarray(term int) []int {
	farr := make([]int, term)
	farr[0], farr[1] = 1, 1
	for i := 2; i < term; i++ {
		farr[i] = farr[i-1] + farr[i-2]
	}
	return farr
}
