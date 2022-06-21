package main

import "fmt"

func main() {
	emp3 := struct {
		firstName, lastName string
		age, salary         int
	}{
		firstName: "Andrea",
		lastName:  "Nikola",
		age:       31,
		salary:    5000,
	}
	fmt.Println(emp3)
}
