package func_05_test

import "testing"

func fib() func() int {
	a, b := 1, 1
	return func() int {
		a, b = b, a+b
		return b
	}
}
func TestFibonacciClosure(t *testing.T) {
	f := fib()
	for i := 0; i <= 9; i++ {
		t.Log(i, f())
	}
}
