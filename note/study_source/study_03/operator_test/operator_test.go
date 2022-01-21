package operator_test

import "testing"

const (
	Readable = 1 << iota
	Writeable
	Executable
)

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 4, 5}
	//c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b)
	//t.Log(a == c)  // 编译错误
	t.Log(a == d)
}

func TestBitClear(t *testing.T) {
	a := 7             // 0111  // 位运算 // 可读可写可执行
	a = a &^ Readable  // 按位清零，去除可读
	a = a &^ Writeable // 按位清零 去除可写
	t.Log(a&Readable == Readable, a&Writeable == Writeable, a&Executable == Executable)
}
