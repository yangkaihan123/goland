package f_test

import "testing"

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func isGreater(x, y int) bool {
	if x > y {
		return true
	}
	return false
}

func TestFn(t *testing.T) {
	a := 22
	b := 20
	t.Log(isGreater(a, b))
	t.Log(Abs(b))
}

//使用简短方式 := 声明的变量的作用域只存在于 if 结构中（在 if 结构的大括号之间，如果使用 if-else 结构则在 else 代码块中变量也会存在）。
//如果变量在 if 结构之前就已经存在，那么在 if 结构中，该变量原来的值会被隐藏。
//最简单的解决方案就是不要在初始化语句中声明变量
