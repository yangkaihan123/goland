package constant_test

import "testing"

const (
	Monday = iota + 1 // iota 默认值为0
	Tuesday
	Wednesday
	thuesday
	Friday
)

const (
	Readable = 1 << iota
	Writeable
	Executable
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, thuesday)
}

func TestConstantTry_01(t *testing.T) {
	a := 1 // 0001  // 位运算
	t.Log(a&Readable == Readable, a&Writeable == Writeable, a&Executable == Executable)
}
