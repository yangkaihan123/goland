package CrashNilPointer_test

import "testing"

func TestPoint(t *testing.T) {
	var i1 = 5
	t.Logf("An integer: %d, it's loaction in memory：%p\n", i1, &i1)

	var intP *int
	intP = &i1 //这里intP指向i1，intP存储了i1的内存地址（指针） 它指向了 i1 的位置，它引用了变量 i1
	t.Logf("An integer: %p, it's loaction in memory：%d\n", intP, *intP)
	//符号 * 可以放在一个指针前，如 *intP，那么它将得到这个指针指向地址上所存储的值；这被称为反引用（或者内容或者间接引用）操作符；另一种说法是指针转移。
}
