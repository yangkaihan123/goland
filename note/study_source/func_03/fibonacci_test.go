package func_03_test

import "testing"

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}
func TestFibonacci(t *testing.T) {
	result := 0
	for i := 0; i <= 10; i++ {
		result = fibonacci(i)
		t.Logf("fibonacci(%d) is: %d\n", i, result)
	}
}
