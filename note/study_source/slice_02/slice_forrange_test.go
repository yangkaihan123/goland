package slice_02

import "testing"

func TestSliceForRange(t *testing.T) {
	var slice1 = make([]int, 128)

	slice1[0] = 1
	slice1[1] = 2
	slice1[2] = 3
	slice1[3] = 4

	for ix, value := range slice1 {
		t.Logf("Slice at %d is %d\n", ix, value)
	}
}
