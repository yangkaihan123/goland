package main

import (
	"fmt"
	"math"
)

func main() {
	x := 121
	y := isPalindrome(x)
	fmt.Println("x is ", y)
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}
	temp := x
	round := 0
	for {
		temp = temp / 10
		round++
		if temp == 0 {
			break
		}
	}
	for round > 0 {
		suffix := x % 10
		preffix := x / int(math.Pow10(round-1))
		if suffix != preffix {
			return false
		}
		x = (x - preffix*int(math.Pow10(round-1))) / 10
		round = round - 2
	}
	return true
}
