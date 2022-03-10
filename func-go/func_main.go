package main

import "fmt"

func test(fn func() int) int {
	return fn()
}

//定义函数类型
type FormatFunc func(s string, x, y int) string

func format(fn FormatFunc, s string, x, y int) string {
	return fn(s, x, y)
}
func main() {
	s1 := test(func() int {
		return 100
	}) // 直接将匿名函数做参数

	s2 := format(func(s string, x, y int) string {
		return fmt.Sprintf(s, x, y)
	}, "%d,%d", 10, 20)

	println(s1, s2)
}
