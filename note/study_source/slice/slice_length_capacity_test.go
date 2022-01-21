package slice

import (
	"fmt"
	"testing"
)

//切片长度与容量相等的情况
func TestSliceLengthCapacity(t *testing.T) {
	numbers := []int{0, 1, 2}
	printSlice(numbers)
	//输出为 len=3 cap=3 slice=[0 1 2]
	/*通过append给切片numbers增加信息，如果当前切片长度和容量相等len=cap，增加信息的长度小于等于原来的长度,
	那么切片的长度变为相加之和，容量变为原来的2倍*/
	numbers = append(numbers, 10, 5, 6)
	printSlice(numbers)
	//输出为 len=6 cap=6 slice=[0 1 2 10 5 6]
	/*通过append 给numbers增加信息，如果当前切片A的长度与容量相等，增加信息B的长度大于切片A那么切片的长度变为相加之和，容量变为：
	B长度+A长度+(B长度-A长度)%2
	*/
	numbers = append(numbers, 12, 13, 15, 16, 17, 18, 19, 20, 21, 22, 23)
	printSlice(numbers)
	//输出为：len=17 cap=18 slice=[0 1 2 10 5 6 12 13 15 16 17 18 19 20 21 22 23]
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
