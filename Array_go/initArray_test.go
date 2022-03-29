package Array_go

import (
	"fmt"
	"testing"
)

func TestInitArray_01(t *testing.T) {
	var testArray [3]int
	var numArray = [3]int{1, 2}
	var cityArray = [3]string{"上海", "北京", "深圳"}
	fmt.Println(testArray, numArray, cityArray)
}

func TestInitArray_02(t *testing.T) {
	var testArray_02 [3]int
	var numArray_02 = [...]int{1, 2}
	var cityArray_02 = [...]string{"上海", "北京", "深圳"}
	fmt.Println(testArray_02)
	fmt.Println(numArray_02)
	fmt.Printf("type of numArray_02 is %T\n", numArray_02)
	fmt.Println(cityArray_02)
	fmt.Printf("type of cityArray_02 is %T\n", cityArray_02)
}

func TestInitArray_03(t *testing.T) {
	a := [...]int{1: 1, 3: 5}
	fmt.Println(a)
	fmt.Printf("type of a is %T\n", a)
}

//遍历数组的两种方式
func TestRangeArray(t *testing.T) {
	var a = [...]string{"上海", "北京", "深圳"}
	// 方法1 ：for循环遍历数组
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	// 方法2 ： for range遍历
	for _, value := range a {
		fmt.Println(value)
	}
}
