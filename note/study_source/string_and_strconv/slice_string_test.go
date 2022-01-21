package string_conversion_test

import (
	"strings"
	"testing"
)

func TestFields(t *testing.T) {
	a := "abcd"
	a1 := strings.Fields(a)
	t.Logf("EXAMPLE in slice is %v\n", a1)
	for _, val := range a1 {
		t.Logf("%s", val)
	}
}

func TestSplit(t *testing.T) {
	a := "GO1|The ABC of Go|25"
	a1 := strings.Split(a, "|")
	t.Logf("the slice is %v\n", a1)
	for _, val := range a1 {
		t.Logf("%s - ", val)
	}
	str1 := strings.Join(a1, ";")
	t.Logf("joined by ; : %s\n", str1)
}
