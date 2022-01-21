package func_02_test

import (
	"fmt"
	"testing"
)

func TestDefer(t *testing.T) {
	function1()
}

func function1() {
	fmt.Printf("In func1 at the top\n")
	defer function2()
	fmt.Printf("in func1 at the bottom!\n")
}

func function2() {
	fmt.Printf("func2: Deferred until the end of the calling func1!")
}

func TestDefer2(t *testing.T) {
	i := 2
	defer t.Logf("the value is : %d\n", i)
	i++
}

func TestDefer3(t *testing.T) {
	for i := 0; i < 6; i++ {
		defer t.Logf("the value is %d\n", i)
	}
}
