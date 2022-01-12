package func02_test

import (
	"fmt"
	"testing"
)

func TestFunc02(t *testing.T) {
	t.Logf("Multiply 2 * 5 * 6 = %d\n", MultiPly3Nums(2, 5, 6))
	// var i1 int = MultiPly3Nums(2, 5, 6)
	// fmt.Printf("MultiPly 2 * 5 * 6 = %d\n", i1)
}

func MultiPly3Nums(a int, b int, c int) int {
	// var product int = a * b *c
	// return product
	return a * b * c
}

// 这里注释信息算是另一种写法，引入了一个本地变量

//命名的返回值
var (
	num   int = 10
	numx2 int
	numx3 int
)

func TestFunctionReturn(t *testing.T) {
	numx2, numx3 = getX2AndX3(num)
	PrintValues()
	numx2, numx3 = getx2andx32(num)
	PrintValues()
}

func getx2andx32(input int) (x2 int, x3 int) {
	x2 = 2 * input
	x3 = 3 * input
	return
}

func PrintValues() {
	fmt.Printf("num = %d, 2x num = %d, 3x num = %d\n", num, numx2, numx3)
}

func getX2AndX3(input int) (int, int) {
	return 2 * input, 3 * input
}
