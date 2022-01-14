package func_02_test

import (
	"io"
	"log"
	"testing"
)

func func1(s string) (n int, err error) {
	defer func() {
		log.Printf("fun1(%q) = %d, %v", s, n, err)
	}()
	return 7, io.EOF
}
func TestDeferLogValues(t *testing.T) {
	func1("Go")
}
