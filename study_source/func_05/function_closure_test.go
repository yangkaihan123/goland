package func_05_test

import "testing"

func TestFunctionClosure(t *testing.T) {
	var f = Adder2()
	t.Log(f(1), " - ")
	t.Log(f(28), " - ")
	t.Log(f(300))
}
func Adder2() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}
