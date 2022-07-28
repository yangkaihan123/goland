package main

import "fmt"

func main() {
	x := min(1, 3, 2, 0)
	fmt.Println("The min is: %d\n", x)

}

func min(s ...int) int {
	if len(s) == 0 {
		return 0
	}
	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}
