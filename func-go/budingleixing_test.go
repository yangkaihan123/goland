package main

import (
	"fmt"
	"testing"
)

func test_02(s string, n ...int) string {
	var x int
	for _, i := range n {
		x += i
	}
	return fmt.Sprintf(s, x)
}
func Test_01(t *testing.T) {
	println(test_02("sum: %d", 1, 2, 5))
}

// 当使用slice 对象作为可变参数传参时，必须展开（slice...)

func Test_03(t *testing.T) {
	s := []int{1, 2, 3}
	res := test_02("sum: %d", s...) // slice... 展开slice
	println(res)
}
