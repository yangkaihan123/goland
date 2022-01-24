package slice_02

import "testing"

func TestForRange02(t *testing.T) {
	seasons := []string{"SPRING", "SUMMER", "AUTUMN", "WINTER"}

	for ix, season := range seasons {
		t.Logf("Season %d is %s\n", ix, season)
	}
}

func TestF(t *testing.T) {
	var a = make([]string, 4)
	a[0] = "1"
	a[1] = "2"
	a[2] = "3"
	a[3] = "4"
	for i, a1 := range a {
		t.Logf("the ix is %d, the a is %s\n", i, a1)
	}
}
