package string_conversion_test

import (
	"strconv"
	"testing"
)

func TestConversion(t *testing.T) {
	orig := "abb"
	var an int
	var newS string
	t.Logf("the size of ints is: %d\n", strconv.IntSize)
	an, _ = strconv.Atoi(orig)
	t.Logf("the integer is: %d\n", an)
	an += 5
	newS = strconv.Itoa(an)
	t.Logf("the new string is %s\n", newS)
}

func TestConversion2(t *testing.T) {
	var a string = "ABC"
	var newS string
	t.Logf("sdasda is %d\n", strconv.IntSize)
	an, err := strconv.Atoi(a)
	if err != nil {
		t.Logf("a %s is not an integer -- exiting with error\n", a)
		return
	}
	t.Logf("the integer is %d\n", an)
	an += 5
	newS = strconv.Itoa(an)
	t.Logf("the new string is %s\n", newS)
}
