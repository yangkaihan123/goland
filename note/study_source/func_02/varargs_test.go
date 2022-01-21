package func_02_test

import (
	"fmt"
	"testing"
)

func TestVarargs(t *testing.T) {
	printInts()
	t.Log()
	printInts(2, 3)
	t.Log()
	printInts(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
}

func printInts(list ...int) {
	//for _, v := range list {
	//	fmt.Printf("%d\n", v)
	//}
	for i := 0; i < len(list); i++ {
		fmt.Printf("%d\n", list[i])
	}
}
