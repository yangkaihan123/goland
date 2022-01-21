package func_03_test

import (
	"fmt"
	"testing"
)

func TestRecursive10To1(t *testing.T) {
	printrec(1)
}

func printrec(i int) {
	if i > 10 {
		return
	}
	printrec(i + 1)
	fmt.Printf("%d\n", i)
}
