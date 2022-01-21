package CrashNilPointer_test

import "testing"

func TestCrashPointer(t *testing.T) {
	var p *int = nil
	*p = 0
	t.Log(*p)
	t.Log(p)
}

//panic: runtime error: invalid memory address or nil pointer dereference [recovered]
//panic: runtime error: invalid memory address or nil pointer dereference
//[signal 0xc0000005 code=0x1 addr=0x0 pc=0x503f1f]
