package arr

import "testing"

func TestArrayPointer(t *testing.T) {
	var arr1 [5]int

	for i := 0; i < len(arr1); i++ {
		arr1[1] = i * 2
	}
	arr2 := arr1
	arr2[2] = 100

	for i := 0; i < len(arr2); i++ {
		t.Logf("Array arr1 at index %d is %d\n", i, arr1[i])
	}
	t.Log()
	for i := 0; i < len(arr2); i++ {
		t.Logf("Array arr2 at index %d is %d\n", i, arr2[i])
	}
}
