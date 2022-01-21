package func_04_test

import (
	"strings"
	"testing"
)

func TestStringsMap(t *testing.T) {
	asciiOnly := func(c rune) rune {
		if c > 127 {
			return ' '
		}
		return c
	}
	t.Log(strings.Map(asciiOnly, "Jérôme Österreich"))
}
