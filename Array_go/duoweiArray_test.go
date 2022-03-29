package Array_go

import (
	"fmt"
	"testing"
)

var a = [3][2]string{
	{"北京", "上海"},
	{"广州", "深圳"},
	{"成都", "重庆"},
}

//二维数组的定义
func TestTwoArray(t *testing.T) {
	fmt.Println(a)
	fmt.Println(a[2][1]) // 索引从0开始
}

//二维数组的遍历
func TestRangeTwoArry(t *testing.T) {
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s\t", v2)
		}
		fmt.Println()
	}
}

/*
多维数组只支持第一层可以使用 ... 来让编译器推导数组长度 可变数组
*/
// 支持的写法：
var a1 = [...][2]string{
	{"北京", "上海"},
	{"广州", "深圳"},
	{"成都", "重庆"},
}

// 不支持多维数组内层使用 ...
var b = [3][...]string{
	{"北京", "上海"},
	{"广州", "深圳"},
	{"成都", "重庆"},
}

func TestArray03(t *testing.T) {
	fmt.Println(a1)
	fmt.Println(b)
}

// .\duoweiArray_test.go:41:12: use of [...] array outside of array literal
