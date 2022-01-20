package arr

import (
	"fmt"
	"testing"
)

func TestPointerArray2(t *testing.T) {
	for i := 0; i < 3; i++ {
		fp(&[3]int{i, i * i, i * i * i})
	}
}
func fp(a *[3]int) {
	fmt.Println(a)
}
