package func_03_test

import "testing"

func TestDigui(t *testing.T) {
	t.Logf("%d is even: is %t\n", 16, even(16))
}
func even(nr int) bool {
	if nr == 0 {
		return true
	}
	return odd(RevSign(nr) - 1)
}
func RevSign(nr int) int {
	if nr < 0 {
		return -nr
	}
	return nr
}
func odd(nr int) bool {
	if nr == 0 {
		return false
	}
	return even(RevSign(nr) - 1)
}
