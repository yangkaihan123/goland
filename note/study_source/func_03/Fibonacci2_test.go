package func_03_test

import "testing"

func fibonacci2(n int) (val, pos int) {
	if n <= 1 {
		val = 1
	} else {
		v1, _ := fibonacci2(n - 1)
		v2, _ := fibonacci2(n - 2)
		val = v1 + v2
	}
	pos = n
	return
}

func TestFibonacci02(t *testing.T) {
	pos := 5
	res, pos := fibonacci2(pos)
	t.Logf("the %d-th Fibonacci number is %d\n", pos, res)
	pos = 10
	res, pos = fibonacci2(pos)
	t.Logf("the %d-th Fibonacci number is %d\n", pos, res)
}
