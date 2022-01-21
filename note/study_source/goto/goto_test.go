package goto_test

import "testing"

//for、switch 或 select 语句都可以配合标签（label）形式的标识符使用，即某一行第一个以冒号（:）结尾的单词（gofmt 会将后续代码自动移至下一行）
func TestLabel(t *testing.T) {
LABEL1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				continue LABEL1
			}
			t.Logf("i is %d, and j is %d\n", i, j)
		}
	}
}

//continue指向LABEL1，当循环执行到j==4的时候会跳转到LABEL1的位置
//所以当j==4和j==5的时候没有输出
//标签的作用对象为外部循环，因此i会直接变成下一个循环的值，此时j的值会重置为0（j的初始值）
//如果直接用break则不会只退出内层循环，而是直接退出外层循环

func TestGoto(t *testing.T) {
	i := 0
HERE:
	t.Log(i)
	i++
	if i == 5 {
		return
	}
	goto HERE
}

//使用goto和标签来模拟循环
//使用标签和 goto 语句是不被鼓励的：它们会很快导致非常糟糕的程序设计，而且总有更加可读的替代方案来实现相同的需求。
//func TestGoto2(t *testing.T) {
//	a := 1
//	goto TARGET // compile error
//	b := 9
//TARGET:
//	b += a
//	t.Logf("a is %v *** b is %v", a, b)
//}

//使用 goto，应当只使用正序的标签（标签位于 goto 语句之后），但注意标签和 goto 语句之间不能出现定义新变量的语句，否则会导致编译失败

func TestThink1(t *testing.T) {
	i := 0
	for { //since there are no checks, this is an infinite loop
		if i >= 3 {
			break //break out of this for loop when this condition is met
		}
		t.Logf("the value of i is : %d\n", i)
		i++
	}
	t.Log("A statement just after for loop ")
}

func TestThink2(t *testing.T) {
	for i := 0; i < 7; i++ {
		if i%2 == 0 {
			continue
		}
		t.Logf("Odd: %d\n", i)
	}
}
