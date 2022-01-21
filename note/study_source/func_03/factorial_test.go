package func_03_test

import "testing"

func Factorial(n uint64) (fac uint64) {
	fac = 1
	if n > 0 {
		fac = n * Factorial(n-1)
		return
	}
	return
}
func TestFactorial(t *testing.T) {
	for i := uint64(0); i < uint64(30); i++ {
		t.Logf("Factorial of %d is %d\n", i, Factorial(i))
	}
}

/* unnamed return variables: 返回未命名返回值
func Factorial_02(n uint64) uint64 {
	if n > 0 {
		return n * Factorial(n-1)
	}
	return 1
}
*/
