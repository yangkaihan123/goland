package error_test

import "testing"

func TestIfElse02(t *testing.T) {
	var first int
	var second int

	if first = -10; first <= 0 {
		t.Logf("first is less than 0 %d\n", first)
	} else if first > 0 && first < 5 {
		t.Log("first is between 0 and 5\n")
	} else {
		t.Log("first is 5 or greater\n")
	}

	if second = 5; second > 10 {
		t.Log("second is greater 10\n")
	} else {
		t.Log("second is less than 10\n")
	}
}
