package main

import "fmt"

func main() {
	var value int
	var isPresent bool

	map1 := make(map[string]int)
	map1["New Delhi"] = 55
	map1["beijing"] = 20
	map1["washington"] = 25
	value, isPresent = map1["beijing"]
	if isPresent {
		fmt.Printf("The value of \"beijing\" in map1 is %d\n", value)
	} else {
		fmt.Printf("map1 does not contain beijing")
	}

	value, isPresent = map1["Paris"]
	fmt.Printf("Is \"Paris\" in map1 ? %t\n", isPresent)
	fmt.Printf("value is %d\n", value)
}
