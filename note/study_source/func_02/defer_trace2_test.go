package func_02_test

import (
	"fmt"
	"testing"
)

func trace_new(s string) string {
	fmt.Println("entering:", s)
	return s
}
func un(s string) {
	fmt.Println("leaving:", s)
}
func a() {
	defer un(trace_new("a"))
	fmt.Println("in a")
}
func b() {
	defer un(trace_new("b"))
	fmt.Println("in b")
	a()
}
func TestDefertrace(t *testing.T) {
	b()
}
