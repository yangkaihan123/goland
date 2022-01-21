package arr

import (
	"fmt"
	"testing"
)

func f(a [3]int) {
	fmt.Println(a)
}
func f2(a *[3]int) {
	fmt.Println(a)
}
func TestPointerArray(t *testing.T) {
	var ar [3]int
	f(ar)
	f2(&ar)
}
