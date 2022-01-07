package string_conversion_test

import (
	"strings"
	"testing"
)

func TestNewReader(t *testing.T) {
	a := "test new reader"
	a1 := strings.NewReader(a)
	t.Logf("the strings is %s\n", a)
	t.Logf("the pointer is %v\n", a1)
	t.Logf("the pointer value is %v\n", *a1)
}
