package main

import "fmt"

func main() {
	emp6 := Employee{"Sam", "Andrew", 55, 6000}
	fmt.Println(emp6.age)
	fmt.Println(emp6.firstName)
	fmt.Println(emp6.lastName)
	fmt.Println(emp6.salary)
}
