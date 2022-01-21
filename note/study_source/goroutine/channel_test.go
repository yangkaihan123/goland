package lock_test

import (
	"testing"
)

func TestChannel(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	go func() {
		for {
			i, ok := <-ch1
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()
	for i := range ch2 {
		t.Log(i)
	}
}

func TestSelect(t *testing.T) {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			t.Log(x)
		case ch <- i:
		}
	}
}
