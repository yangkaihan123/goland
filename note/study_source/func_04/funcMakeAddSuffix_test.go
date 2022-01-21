package func_04_test

import (
	"strings"
	"testing"
)

func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
func TestMakeAddSuffix(t *testing.T) {
	addBmp := MakeAddSuffix(".bmp")
	addJpeg := MakeAddSuffix(".jpeg")
	t.Log(addBmp("file"))
	t.Log(addJpeg("file"))
}
