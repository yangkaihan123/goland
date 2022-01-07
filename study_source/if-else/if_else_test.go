package error_test

import (
	"runtime"
	"testing"
)

func TestIfElse(t *testing.T) {
	bool1 := true
	if bool1 {
		t.Log("the value is true\n")
	} else {
		t.Log("the value is false\n")
	}
}

func TestCondition(t *testing.T) {
	a := ""
	//if a == " " {
	//	t.Log("str is null")
	//} else {
	//	t.Log("str is not nul",a)
	//}
	if len(a) != 0 {
		t.Log("not null")
	} else {
		t.Log("is null")
	} //两种写法
}

func TestRuntimeOS(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Log("windows")
	} else if runtime.GOOS == "linux" {
		t.Log("linux")
	} else {
		t.Log("other")
	}
}
