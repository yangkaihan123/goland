package func_04_test

import "testing"

func TestReturnDefer(t *testing.T) {
	t.Log(f2())
}
func f2() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}
