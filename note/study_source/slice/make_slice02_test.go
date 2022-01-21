package slice

import "testing"

func TestMakeSlice02(t *testing.T) {
	s := make([]byte, 5)
	t.Logf("the slice length is %d and capacity is %d\n", len(s), cap(s))
	//the slice length is 5 and capacity is 5
	s = s[2:4]
	t.Logf("the slice length is %d and capacity is %d\n", len(s), cap(s))
	//the slice length is 2 and capacity is 3
}
