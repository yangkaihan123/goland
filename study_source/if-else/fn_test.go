package f_test

import "testing"

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func isGreater(x, y int) bool {
	if x > y {
		return true
	}
	return false
}

func TestFn(t *testing.T) {
	x := 10
	y := 20
	t.Log(isGreater(x, y))
	t.Log(Abs(x))
}
