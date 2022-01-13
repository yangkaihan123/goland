package func03_test

import "testing"

func TestBlank(t *testing.T) {
	var (
		i1 int
		f1 float32
	)
	i1, _, f1 = ThreeValues()
	t.Logf("the int is %d, the float is %f \n", i1, f1)
}

func ThreeValues() (int, int, float32) {
	return 5, 6, 7.5
}

func TestReturnBlank(t *testing.T) {
	var (
		min int
	)
	min, _ = MinMax(78, 65)
	t.Logf("Minmub is %d, Maxmub is _", min)
}

func MinMax(i int, i2 int) (min int, max int) {
	if i < i2 {
		min = i
		max = i2
	} else {
		min = i2
		max = i
	}
	return min, max
}
