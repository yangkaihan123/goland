package CrashNilPointer_test

import "testing"

func TestStringPointer(t *testing.T) {
	s := "good bye"
	var p *string = &s
	*p = "cc"
	t.Logf("p's pointer: %p\n", p)
	t.Logf("*p is : %s\n", *p)
	t.Logf("s is : %s\n", s)
	t.Logf("s's pointer: %p\n", &s)
}

//string_pointer_test.go:9: p's pointer: 0xc000098520
//string_pointer_test.go:10: *p is : cc
//string_pointer_test.go:11: s is : cc
//string_pointer_test.go:12: s's pointer: 0xc000098520
//通过对*p的赋值更改了s，这里给*p赋值，相当于直接给s的指针赋值了
//不能获取常量的地址，不能获取字面量的地址
// const i = 5  ptr := &i error: cannot take the address of i
// ptr2 := &10  error: cannot take the address of 10
