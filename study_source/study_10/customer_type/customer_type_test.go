package customer_type

import (
	"fmt"
	"testing"
	"time"
)

type Init func(op int) int // 把方法自定义一个类型

func timeSpent(inner Init) Init { // 这个函数的参数是一个方法
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spents:", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFn(t *testing.T) {
	//忽略其中一个返回值可以使用 _ 下划线代替
	//a, _ := returnMultiValues()
	//t.Log(a)
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}
