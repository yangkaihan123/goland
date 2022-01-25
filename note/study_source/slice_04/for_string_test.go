package slice_04

import "testing"

func TestForString(t *testing.T) {
	s := "\u00ff\u754c"
	for i, c := range s {
		t.Logf("%d:%c", i, c)
	}
}
