package func03_test

import "testing"

func TestReturn(t *testing.T) {
	sum, prod, diff := SumProductDiff(5, 6)
	t.Logf("the sum is %d\n, the prod is %d\n, the diff is %d\n", sum, prod, diff)

	sum2, prod2, diff2 := SumProductDiff2(6, 7)
	t.Log(sum2, prod2, diff2)
}
func SumProductDiff(a, b int) (x, y, z int) {
	x = a + b
	y = a * b
	z = a - b
	return x, y, z
}
func SumProductDiff2(a, b int) (int, int, int) {
	return a + b, a * b, a - b
}
