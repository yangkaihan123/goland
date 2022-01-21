package lock_test

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

var wg sync.WaitGroup

func Hello(i int) {
	defer wg.Done() // goroutine 结束就登记-1
	fmt.Println("Hello Goroutine!")
}
func TestRoutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		wg.Add(1) // 启动一个routine就登记+1
		go Hello(i)
	}
	wg.Wait() //等待全部登记的routine都结束
}

func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A", i)
	}
}
func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B", i)
	}
}
func TestGPM(t *testing.T) {
	runtime.GOMAXPROCS(2)
	go a()
	go b()
	time.Sleep(time.Second)
}
