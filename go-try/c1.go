package main

// #include <stdlib.h>
import "C"

func Random() int {
	return int(C.random())
}
func seed(i int) {
	C.srandom(C.uint(i))
}
