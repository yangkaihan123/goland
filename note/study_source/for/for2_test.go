package for2_test

import "testing"

func TestBreak(t *testing.T) {
	i := 5
	for {
		i += 1
		t.Logf("the value i is now: %d\n", i)
		if i > 1000 {
			break
		}
	}
}

func TestBreak2(t *testing.T) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 10; j++ {
			if j > 5 {
				break //这里break只跳出最内层循环
			}
			t.Log(j)
		}
		t.Log(" ")
	}
}

func TestContinue(t *testing.T) {
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue //忽略剩余的循环体直接进入下次循环过程，但是不是无条件执行下一次循环，执行前依旧遵循循环的判断条件
		} //continue只能用于for循环中
		t.Log(i)
		t.Log(" ")
	}
}
