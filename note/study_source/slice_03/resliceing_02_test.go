package slice_03

import "testing"

func TestReslicing2(t *testing.T) {
	s := make([]int, 5, 5)
	t.Logf("the slice s length is %d, the cap is %d\n", len(s), cap(s))
	s = s[2:2]
	t.Logf("the slice s length is %d, the cap is %d\n", len(s), cap(s))
	s = s[2:3]
	t.Logf("the slice s length is %d, the cap is %d\n", len(s), cap(s))
}
