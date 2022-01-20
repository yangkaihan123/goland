package slice

import "testing"

func TestArraySlice(t *testing.T) {
	var arr1 [6]int
	var slice1 []int = arr1[2:5] // item at index 5 not included

	//load the array with integers: 0,1,2,3,4,5
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}

	//print the slice
	for i := 0; i < len(slice1); i++ {
		t.Logf("Slice at %d is %d\n", i, slice1[i])
	}

	t.Logf("the length of arr1 is %d\n", len(arr1))
	t.Logf("the length of slice1 is %d\n", len(slice1))
	t.Logf("the capacity of slice1 is %d\n", cap(slice1))

	//grow the slice
	slice1 = slice1[0:4]
	for i := 0; i < len(slice1); i++ {
		t.Logf("Slice at %d is %d\n", i, slice1[i])
	}
	t.Logf("the length of slice1 is %d\n", len(slice1))
	t.Logf("the capacity of slice1 is %d\n", cap(slice1))

	// grow the slice beyond capacity
	//slice1 = slice1[0:7 ] // panic: runtime error: slice bound out of range
}
