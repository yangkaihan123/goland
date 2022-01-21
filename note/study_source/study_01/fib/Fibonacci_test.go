package fib

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	//var a int = 1
	//var b int = 1

	//var (
	//	a int = 1
	//	b int = 2
	//)
	// 变量命名，赋值的三种方式
	a := 1
	b := 2
	t.Log(a)
	for i := 0; i < 5; i++ {
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
}

func TestExchange(t *testing.T) {
	a := 1
	b := 2
	a, b = b, a
	t.Log(a, b)
}
