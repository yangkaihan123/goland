package arr

import "testing"

func TestForArray_02(t *testing.T) {
	var a [15]int
	for i := 0; i < len(a); i++ {
		a[i] = i
	}
	t.Log(a)
}
