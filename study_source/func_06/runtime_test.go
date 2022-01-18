package func_06_test

import (
	"testing"
	"time"
)

func TestRuntime(t *testing.T) {
	result := 0
	start := time.Now()
	for i := 0; i <= 25; i++ {
		result = fibonacci(i)
		t.Logf("fibonacci(%d) is %d\n", i, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	t.Logf("fibonacci took this amount of time: %s\n", delta)
}

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}
