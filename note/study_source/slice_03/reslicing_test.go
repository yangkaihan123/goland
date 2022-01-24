package slice_03

import "testing"

func TestReslicing(t *testing.T) {
	slice1 := make([]int, 0, 10)
	//load the slice, cap(slice1) is 10
	for i := 0; i < cap(slice1); i++ {
		slice1 = slice1[0 : i+1]
		slice1[i] = i
		t.Logf("the length of slice is %d\n", len(slice1))
	}
	//print the slice
	for i, a := range slice1 {
		t.Logf("Slice at %d is %d\n", i, a)
	}
}
