package slice_02

import (
	"testing"
)

func TestForRange(t *testing.T) {
	items := [...]int{10, 20, 30, 40, 50}
	for _, item := range items {
		item *= 2
		t.Logf(string(items[item]))
	}
} // 无法编译
//不理解这里为什么是string类型
