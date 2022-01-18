package func_06_test

import (
	"testing"
	"time"
)

const LIM = 41

var fibs [LIM]uint64

func TestFibonacciMem(t *testing.T) {
	var result uint64 = 0
	start := time.Now()
	for i := 0; i < LIM; i++ {
		result = fibonacci02(i)
		t.Logf("fibonacci(%d) is %d\n", i, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	t.Logf("fibonacci took this amount of time is %s\n", delta)
}

func fibonacci02(n int) (res uint64) {
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci02(n-1) + fibonacci02(n-2)
	}
	fibs[n] = res
	return
}
