package slice_02

import (
	"math"
	"testing"
)

func TestMinMax(t *testing.T) {
	sl1 := []int{78, 34, 643, 12, 90, 492, 13, 2}
	max := maxSlice(sl1)
	t.Logf("the slice max is %d\n", max)
	min := minSlice(sl1)
	t.Logf("the slice min is %d\n", min)
}

func minSlice(sl []int) (min int) {
	min = math.MaxInt32
	for _, b := range sl {
		if b < min {
			min = b
		}
	}
	return
}

func maxSlice(sl []int) (max int) {
	for _, a := range sl {
		if a > max {
			max = a
		}
	}
	return
}
