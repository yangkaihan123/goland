package error_test

import (
	"strconv"
	"testing"
)

func TestError(t *testing.T) {
	var orig string = "ABC"
	var newS string
	t.Logf("the size of ints is %d\n", strconv.IntSize)
	an, err := strconv.Atoi(orig)
	if err != nil {
		t.Logf("orig %s is not an integer - exiting with error\n", orig)
		return
	}
	t.Logf("the integer is %d\n", an)
	an += 5
	newS = strconv.Itoa(an)
	t.Logf("the new string is %s\n", newS)
}
