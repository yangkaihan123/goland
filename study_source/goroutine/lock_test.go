package lock_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var x int64
var wg2 sync.WaitGroup
var lock sync.Mutex
var rwlock sync.RWMutex

func add() {
	for i := 0; i < 5000; i++ {
		lock.Lock()
		x += 1
		lock.Unlock()
	}
	wg2.Done()
}

func TestLock(t *testing.T) {
	wg2.Add(2)
	go add()
	go add()
	wg2.Wait()
	t.Log(x)
}

func write() {
	rwlock.Lock()
	x += 1
	time.Sleep(10 * time.Millisecond)
	rwlock.Unlock()
	wg2.Done()
}

func read() {
	rwlock.Lock()
	time.Sleep(time.Millisecond)
	rwlock.Unlock()
	wg2.Done()
}

func TestWriteReadLock(t *testing.T) {
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg2.Add(1)
		go write()
	}
	for i := 0; i < 1000; i++ {
		wg2.Add(1)
		go read()
	}
	wg2.Wait()
	end := time.Now()
	t.Log(end.Sub(start))
}

func hello() {
	defer wg2.Done()
	fmt.Println("hello goroutine!")
}
func TestSync(t *testing.T) {
	wg2.Add(1)
	go hello()
	t.Log("main goroutine done!")
	wg2.Wait()
}
