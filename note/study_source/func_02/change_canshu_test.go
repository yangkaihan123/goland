package func_02_test

import "testing"

func TestCanChange(t *testing.T) {
	x := min(1, 3, 2, 0)
	t.Logf("the minfunc is %d\n", x)
	b := []int{7, 9, 3, 5, 1}
	a := min(b...)
	t.Logf("the minfunc is %d\n", a)
}

func min(s ...int) int {
	if len(s) == 0 {
		return 0
	}
	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}
