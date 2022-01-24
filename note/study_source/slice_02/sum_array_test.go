package slice_02

import "testing"

func TestSumArray(t *testing.T) {
	var a = []float32{1.0, 2.0, 3.0, 4.0}
	t.Logf("the sum of array is : %f\n", Sum(a))
	var b = []int{1, 2, 3, 4, 5}
	sum, average := SumAndAverage(b)
	t.Logf("the sum of array is %d, and the average is %f", sum, average)
}

func SumAndAverage(a []int) (int, float32) {
	sum := 0
	for _, item := range a {
		sum += item
	}
	return sum, float32(sum / len(a))
}

func Sum(a []float32) (sum float32) {
	for _, item := range a {
		sum += item
	}
	return
}
