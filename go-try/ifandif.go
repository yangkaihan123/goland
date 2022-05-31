package main

import "fmt"

func main() {
	a := 5
	b := 2
	if a+b > 10 {
		if b-a < 0 {
			fmt.Println(a)
		}
	}
}
