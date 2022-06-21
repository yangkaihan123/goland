package main

import "fmt"

type Employee struct {
	firstName, lastName string
	age, salary         int
}

func main() {
	emp1 := Employee{
		firstName: "Sam",
		age:       25,
		salary:    500,
		lastName:  "Anderson",
	}

	emp2 := Employee{"Thomas", "Paul", 29, 800}

	fmt.Println(emp1, emp2)
}
