package slice_03

import "testing"

func TestCopyAppend(t *testing.T) {
	s1From := []int{1, 2, 3}
	s1To := make([]int, 10)

	n := copy(s1To, s1From)
	t.Log(s1To)
	t.Logf("Copied %d elements\n", n) // n == 3

	sl := []int{1, 2, 3}
	sl = append(sl, 4, 5, 6)
	t.Log(sl)
}
