package func03_test

import "testing"

func TestOutsideValue(t *testing.T) {
	n := 0
	reply := &n //n的指针地址
	t.Log("the n's pointer is : ", reply)
	Multiply(10, 5, reply)
	t.Log("reply is ", *reply)
	t.Log("the *reply is ", &*reply)
}
func Multiply(a, b int, reply *int) {
	*reply = a * b
}
