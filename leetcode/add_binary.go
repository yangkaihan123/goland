package main

import "fmt"

func main() {
	a := "11"
	b := "1"
	c := addBinary(a, b)
	fmt.Println("sum= ", c)
	fmt.Println()
}
func addBinary(a string, b string) string {
	for len(a) > len(b) {
		b = "0" + b
	}
	for len(b) > len(a) {
		a = "0" + a
	}
	p := len(b) - 1
	res := ""
	dight := 0
	for p >= 0 {
		ap := int(a[p] - '0')
		bp := int(b[p] - '0')
		sum := ap + bp + dight
		switch sum {
		case 0:
			res = "0" + res
			dight = 0
		case 1:
			res = "1" + res
			dight = 0
		case 2:
			res = "0" + res
			dight = 1
		case 3:
			res = "1" + res
			dight = 1
		}
		p--
	}
	if dight == 1 {
		return "1" + res
	}
	if res == "" {
		return "0"
	}
	return res
}
