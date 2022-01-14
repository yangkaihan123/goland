package func_04_test

import (
	"fmt"
	"testing"
)

func TestCallBack(t *testing.T) {
	callback(1, Add)
}
func Add(a, b int) {
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
}
func callback(y int, f func(int, int)) {
	f(y, 2) //this becomes Add(1,2)
}
