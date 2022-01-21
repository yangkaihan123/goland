package string_test

import (
	"testing"
	"unsafe"
)

func TestString(t *testing.T) {
	var s string
	t.Log(s)
	s = "Hello"
	t.Log(len(s))
	// s[1] = '3' // string是不可变的byte slice
	s = "\xE4\xB8\xA5"
	//s = "\xE4\xBA\xBB\xFF"
	t.Log(s)
	t.Log(len(s))

	s = "中"
	t.Log(len(s))

	c := []rune(s) // 把字符串转换为slice rune()函数
	t.Log(len(c))
	t.Log("rune size: ", unsafe.Sizeof(c[0])) // rune的长度
	t.Logf("中 unicode %x", c[0])              // Unicode的值
	t.Logf("中 UTF8 %x", s)                    // utf8的长度
}

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range s {
		t.Logf("%[1]c %[1]d", c) // 本质是遍历数组，遍历处出来是rune
		t.Logf("%[1]c %[1]x", c) // 遍历出来是byte值
	}
}
