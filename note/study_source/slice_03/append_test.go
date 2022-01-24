package slice_03

import "testing"

func TestAppend(t *testing.T) {
	a := []byte{1, 2, 3, 4}
	AppendByte(a)
}

func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)

	if n > cap(slice) {
		// if necessary, reallocate allocate double what's needed, for future growth
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}
