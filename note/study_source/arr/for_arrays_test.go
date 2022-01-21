package arr

import "testing"

func TestForArray(t *testing.T) {
	var arr1 [5]int
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i * 2
	}
	for i := 0; i < len(arr1); i++ {
		t.Logf("Array1 at the index %d is %d\n", i, arr1[i])
	}
}

func TestArray2(t *testing.T) {
	a := [...]string{"a", "b", "c", "d"}
	for i := range a {
		t.Log("Array item", i, "is", a[i])
	}
}
