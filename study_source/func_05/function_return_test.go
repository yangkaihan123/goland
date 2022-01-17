package func_05_test

import "testing"

func TestFunctionReturn(t *testing.T) {
	// make an Add2 function, give it a name p2, and call it
	p2 := Add2()
	t.Logf("Call Add2 for 3 gives: %v\n", p2(3))
	//make a special Adder function, a gets value 2:
	TwoAdder := Adder(2)
	t.Logf("The result is %v\n", TwoAdder(3))
}

func Add2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}
func Adder(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}
