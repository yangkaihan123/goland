package fn_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func timeSpent(inner func(op int) int) func(op int) int {
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

//装饰者模式
func TestFn(t *testing.T) {
	a, b := returnMultiValues()
	t.Log(a, b)
	//忽略其中一个返回值可以使用 _ 下划线代替
	//a, _ := returnMultiValues()
	//t.Log(a)
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}

func Sum(ops ...int) int { //可编程参数
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}
func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4))
	t.Log(Sum(1, 2, 3, 4, 5))
}

func Clear() {
	fmt.Println("Clear resources")
}

func TestDefer(t *testing.T) {
	defer Clear()
	fmt.Println("start")
	panic("err")
	//fmt.Println("End") // 不可达代码
}

//执行结果：
//=== RUN   TestDefer
//start 先执行了start
//Clear resources
//--- PASS: TestDefer (0.00s)
//PASS

//defer 可以用来释放资源，解锁
