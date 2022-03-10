package main

import "fmt"

func test_01(x, y int, s string) (int, string) {
	n := x + y
	return n, fmt.Sprintf(s, n)
}

func main() {
	a := 1
	b := 2
	c := "aaa"
	res, _ := test_01(a, b, c)
	println(res)
}
