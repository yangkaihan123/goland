package string_test

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFn(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",") //字符串分割
	for _, part := range parts {
		t.Log(part)
	}
	t.Log(strings.Join(parts, "-")) // 重新连接起来
}

func TestConv(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log("str" + s) //字符串转换
	//t.Log(10+strconv.Atoi("10")) //会编译错误需要判断错误值，返回的是两个值
	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(10 + i) // 这样编译没有问题
	}
}
