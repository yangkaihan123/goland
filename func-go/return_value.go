package main

import "fmt"

func add(a, b int) (c int) {
	c = a + b
	return c
}

func calc(a, b int) (sum, avg int) {
	sum = a + b
	avg = (a + b) / 2
	return
}
func main() {
	var a, b = 1, 2
	c := add(a, b)
	sum, avg := calc(a, b)
	fmt.Println(a, b, c, sum, avg)
}
