package type_test

import "testing"

type MyInt int64

func TestImplicit(t *testing.T) {
	var a int = 1
	var b int64
	//b = a
	b = int64(a)
	var c MyInt
	c = MyInt(b)   // 不能直接c=b 别名类型也不可以隐式转换
	t.Log(a, b, c) //cannot use a (type int) as type int64 in assignment,不支持隐式类型转换
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	//aPtr = aPtr + 1  // 编译错误，指针不能运算
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*") // string初始化值是空字符串
	t.Log(len(s))
	//if s == ""{
	//
	//}
}
