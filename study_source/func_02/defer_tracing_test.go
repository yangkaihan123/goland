package func_02_test

import (
	"fmt"
	"testing"
)

func trace(s string) {
	fmt.Println("entering:", s)
}
func untrace(s string) {
	fmt.Println("leaving:", s)
}
func a() {
	trace("a")
	defer untrace("a")
	fmt.Println("in a")
}
func b() {
	trace("b")
	defer untrace("b")
	fmt.Println("in b")
	a()
}
func TestTrace(t *testing.T) {
	b()
}
