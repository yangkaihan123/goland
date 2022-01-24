package slice_02

import "testing"

var arr []byte = []byte{'a', 'b', 'a', 'a', 'a', 'a', 'c', 'd', 'e', 'f', 'g'}

func TestString(t *testing.T) {
	arru := make([]byte, len(arr))
	ixu := 0
	tmp := byte(0)
	for _, val := range arr {
		if val != tmp {
			arru[ixu] = val
			t.Logf("%c", arru[ixu])
			ixu++
		}
		tmp = val
	}
}
