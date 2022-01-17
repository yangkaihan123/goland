package func_04_test

import (
	"fmt"
	"testing"
)

func TestFv(t *testing.T) {
	fv := func() {
		fmt.Println("hello world")
	}
	fv()
	t.Logf("the type of fv is %T", fv)
}
