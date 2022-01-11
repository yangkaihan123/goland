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
