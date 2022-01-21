package slice

import "testing"

func TestMakeSlice(t *testing.T) {
	var slice1 []int = make([]int, 10)
	//load the array/slice
	for i := 0; i < len(slice1); i++ {
		slice1[i] = 5 * i
	}
	//print the slice:
	for i := 0; i < len(slice1); i++ {
		t.Logf("Slice at %d is %d\n", i, slice1[i])
	}
	t.Logf("\nThe length of slice1 is %d\n", len(slice1))
	t.Logf("The capacity of slice1 is %d\n", cap(slice1))
}
