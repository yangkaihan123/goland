package for_test

import "testing"

func TestFor1(t *testing.T) {
	for i := 2; i < 5; i++ {
		t.Logf("this is the %d iteration\n", i)
	}
}

func TestFor2(t *testing.T) {
	str := "Go is a beautiful language"
	t.Logf("the len of str is %d\n", len(str))
	for i := 0; i < len(str); i++ {
		t.Logf("Character on position %d is %c \n", i, str[i])
	}
	str2 := "中华人民共和国"
	t.Logf("the len of str is %d\n", len(str2))
	for j := 0; j < len(str2); j++ {
		t.Logf("str2 on position %d is %c \n", j, str2[j])
	}
}

//ASCII 编码的字符占用 1 个字节，既每个索引都指向不同的字符，而非 ASCII 编码的字符（占有 2 到 4 个字节）不能单纯地使用索引来判断是否为同一个字符

func TestLoopFor(t *testing.T) {
	for i := 0; i < 15; i++ {
		t.Logf("the count is %d\n", i)
	}

	i := 0
START:
	t.Logf("the count is %d\n", i)
	i++
	if i < 15 {
		goto START
	}
}
