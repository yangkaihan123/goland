package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	var goos string = runtime.GOOS
	fmt.Println(goos)
	path := os.Getenv("GOROOT")
	fmt.Println("path:", path)
}
